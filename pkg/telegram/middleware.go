package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type TauthData struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	AuthDate  int64  `json:"auth_date"`
	Hash      string `json:"hash"`
}

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

		tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretkey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
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

		claims := jwt.MapClaims{
			"source":     source,
			"id":         respweb.Data.ID,
			"first_name": respweb.Data.FirstName,
			"username":   respweb.Data.Username,
			"photo_url":  respweb.Data.PhotoURL,
			"auth_date":  respweb.Data.AuthDate,
			"exp":        time.Now().Add(expiration).Unix(),
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

		currentClaims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid claims type"})
			return
		}

		currentClaims["exp"] = time.Now().Add(expiration).Unix()

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
