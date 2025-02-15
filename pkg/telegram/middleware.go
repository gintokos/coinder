package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// to do benchmark of goroutins or just call f
var timeNow int64

func init() {
	go func() {
		for {
			timeNow = time.Now().Unix()
			time.Sleep(time.Second * 60)
		}
	}()
}

type TauthData struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	AuthDate  int64  `json:"auth_date"`
	Hash      string `json:"hash"`
}

type TClaims struct {
	TauthData
	ExpiredAt int64  `json:"exp"`
	Source    string `json:"source"`
}

func (c TClaims) Valid() error {
	if timeNow > c.ExpiredAt {
		return fmt.Errorf("token expired")
	}
	return nil
}

// sets claims in context with type tClaims fields:
//
//	TauthData
//	ExpiredAt int64  `json:"exp"`
//	Source    string `json:"source"`
func AuthMiddleware(cookieName string, btoken string) gin.HandlerFunc {
	h := sha256.New()
	_, err := h.Write([]byte(btoken))
	if err != nil {
		panic(err)
	}
	secretkey := h.Sum(nil)

	return func(c *gin.Context) {
		tokenString, err := c.Cookie(cookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		claims := &TClaims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretkey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if tkn.Valid {
			c.Set("claims", claims)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
	}
}

// expiration time is for jwt token
// maxage determine how old should be request auth_date
// domain and secure params for cookie
func AuthHandler(btoken string, expiration time.Duration, maxage int64, cookieName string, domain string, secure bool) gin.HandlerFunc {
	maxage = maxage * 60 * 60
	return func(c *gin.Context) {
		source := c.GetHeader("T-Source-H")
		if source != "web" && source != "miniapp" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		var respweb struct {
			Datastr string    `json:"datastr"`
			Data    TauthData `json:"data"`
		}

		err := c.BindJSON(&respweb)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid init data"})
			return
		}

		var secretkey []byte
		h := sha256.New()
		h.Write([]byte(btoken))
		secretkeyjwt := h.Sum(nil)

		switch source {
		case "web":
			secretkey = secretkeyjwt
		case "miniapp":
			h := hmac.New(sha256.New, []byte("WebAppData"))
			h.Write([]byte(btoken))
			secretkey = h.Sum(nil)
		}

		hm := hmac.New(sha256.New, secretkey)
		hm.Write([]byte(respweb.Datastr))
		calculatedHash := hex.EncodeToString(hm.Sum(nil))

		if calculatedHash != respweb.Data.Hash {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid hash"})
			return
		}

		if respweb.Data.AuthDate+maxage < time.Now().Unix() {
			c.AbortWithStatus(http.StatusRequestTimeout)
			return
		}

		claims := TClaims{
			TauthData: respweb.Data,
			Source:    source,
			ExpiredAt: time.Now().Add(expiration).Unix(),
		}

		tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := tkn.SignedString(secretkeyjwt)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		setCookie(c, cookieName, tokenString, domain, secure, int(maxage))

		c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
	}
}

func RefreshTokenHandler(btoken string, expiration time.Duration, cookieName string, domain string, secure bool) gin.HandlerFunc {
	h := sha256.New()
	_, err := h.Write([]byte(btoken))
	if err != nil {
		panic(err)
	}
	secretkey := h.Sum(nil)

	return func(c *gin.Context) {
		claimsInterface, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no claims found"})
			return
		}

		currentClaims, ok := claimsInterface.(*TClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid claims type"})
			return
		}

		currentClaims.ExpiredAt = time.Now().Add(expiration).Unix()

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, currentClaims)
		tokenString, err := newToken.SignedString(secretkey)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		setCookie(c, cookieName, tokenString, domain, secure, int(expiration.Seconds()))

		c.JSON(http.StatusOK, gin.H{"message": "Token refreshed"})
	}
}

func setCookie(c *gin.Context, name, value, domain string, secure bool, maxAge int) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
	}
	http.SetCookie(c.Writer, cookie)
}

// middleware for using in local if need claims in logic of app and running without some domain for telegram
// sets some default claims and all reqs will be provided with this claims
func TestMiddleware(data TauthData) gin.HandlerFunc {
	return func(c *gin.Context) {
		source := c.GetHeader("T-Source-H")
		if source != "web" && source != "miniapp" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		claims := TClaims{
			TauthData: data,
			Source:    source,
			ExpiredAt: time.Now().Add(time.Hour * 144).Unix(),
		}

		c.Set("claims", claims)
		c.Next()
	}
}
