package parser

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/pkg/sl"
)

const (
	minRequestInterval    = time.Second * 2
	maxtries              = 3
	expectedAmountOfcoins = 11000
	batchSize             = 500
	maxCoinPerReq         = 5000
	baseURL               = "https://pro-api.coinmarketcap.com"
)

type Parser struct {
	httpClient *http.Client
	db         Database

	apikey    string
	timestamp time.Duration
}

type Database interface {
	SaveCoinsForce(coins []models.DBCoin) error
}

// timestamp is time beetwen planed updating database
type Config struct {
	ApiKeyCoinMarketCap string
	Timestamp           time.Duration
	TimeoutForReq       time.Duration
}

// database can be nil if run is not using
func NewParser(cfg Config, db Database) Parser {
	return Parser{
		httpClient: &http.Client{
			Timeout: cfg.TimeoutForReq,
		},
		apikey:    cfg.ApiKeyCoinMarketCap,
		timestamp: cfg.Timestamp,
		db:        db,
	}
}

// method for parser\main.go only
func (p *Parser) Run(ctx context.Context, wg *sync.WaitGroup) error {
	ticker := time.NewTicker(p.timestamp)
	defer ticker.Stop()

	slog.Debug("First parse was started")

	p.parseallCoins()
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			slog.Debug("Cycle work")
			select {
			case <-ticker.C:
				slog.Debug("ticked, starting parsing")
				p.parseallCoins()

			case <-ctx.Done():
				return
			}
		}
	}()

	<-ctx.Done()

	return nil
}

func (p *Parser) parseallCoins() {
	coins := make([]models.ParserCoin, 0, expectedAmountOfcoins)
	page, limit := 0, maxCoinPerReq

	var flag = true
	for flag {
		start := page*limit + 1
		for i := 0; i < maxtries; i++ {
			slog.Info(fmt.Sprintf("Fetching token list without meta %d-%d...", start, start+limit-1))
			cs, err := p.fetchWitoutMeta(limit, start)
			if err == nil {
				coins = append(coins, cs...)
				if len(cs) < maxCoinPerReq {
					flag = false
				}
				break
			}
			slog.Error("error on fetching coins without meta: ", sl.Err(err))
			time.Sleep(minRequestInterval)
		}
		page++
	}

	for i := 0; i < len(coins); i += batchSize {
		end := i + batchSize
		if end > len(coins) {
			end = len(coins)
		}

		slog.Info(fmt.Sprintf("Fetching token list for meta %d-%d...", i, end))
		for k := 0; k < maxtries; k++ {
			err := p.fetchForMeta(coins[i:end])
			if err == nil {
				break
			}
			slog.Error("error on fetching for meta", sl.Err(err))
			time.Sleep(minRequestInterval)
		}
	}

	slog.Info("updating database")
	dbcoins := make([]models.DBCoin, 0, len(coins))
	startTime := time.Now()
	for i := range coins {
		dbcoins = append(dbcoins, models.ToDBcoin(&coins[i]))
	}
	err := p.db.SaveCoinsForce(dbcoins)
	if err != nil {
		slog.Error("error on saving in database:", sl.Err(err))
	}
	elapsed := time.Since(startTime).Milliseconds()
	slog.Debug(fmt.Sprintf("Operation took %d ms", elapsed))

	slog.Info("all tokens was updated")
}

// ids may be nil
func (p *Parser) fetchWitoutMeta(limit, start int, ids ...string) ([]models.ParserCoin, error) {
	var url string
	if len(ids) > 0 {
		idString := strings.Join(ids, ",")
		url = fmt.Sprintf("%s/v1/cryptocurrency/listings/latest?id=%s&convert=USD",
			baseURL, idString)
	} else {
		url = fmt.Sprintf("%s/v1/cryptocurrency/listings/latest?start=%d&limit=%d&convert=USD",
			baseURL, start, limit)

	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error on creating req: %v", err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", p.apikey)
	req.Header.Set("Accept", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error in reading body response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "    ")
		if err != nil {
			return nil, fmt.Errorf("API returned status: %d, raw error: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("API returned error: %s", prettyJSON.String())
	}

	var response struct {
		Data []models.ParserCoin `json:"data"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error on unmarshaling: %v", err)
	}

	return response.Data, nil
}

// modificate slice
func (p *Parser) fetchForMeta(coins []models.ParserCoin) error {
	var builder strings.Builder
	for i, coin := range coins {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(strconv.Itoa(coin.ID))
	}
	ids := builder.String()

	url := fmt.Sprintf("%s/v2/cryptocurrency/info", baseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error on creating req: %v", err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", p.apikey)
	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Add("id", ids)
	req.URL.RawQuery = q.Encode()

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error in reading body response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned error: %s", string(body))
	}

	var response struct {
		Data map[string]models.ParserMetaDataCoin `json:"data"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("error on unmarshaling response: %v", err)
	}

	counter := 0
	for _, v := range response.Data {
		if counter == len(coins) {
			break
		}
		coins[counter].ParserMetaDataCoin = v
		counter++
	}

	return nil
}

// // use withoun run no meta info like logo png
// func (p *Parser) GetListWithoutMeta(ctx context.Context, idList []int) ([]models.Coin, error) {

// }
