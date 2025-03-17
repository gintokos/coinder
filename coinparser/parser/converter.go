package parser

import (
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"github.com/gintokos/coinder/backend/internal/models"
	psModels "github.com/gintokos/coinder/coinparser/models"
)

func ToDBcoin(pc *psModels.ParserCoin) models.DBCoin {
	parseTime := func(timeStr string) *string {
		if timeStr == "" {
			return nil
		}
		if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
			pgTime := t.Format("2006-01-02 15:04:05")
			return &pgTime
		}
		return nil
	}

	joinURLs := func(urls []string) string {
		return strings.Join(urls, ",")
	}

	return models.DBCoin{
		ID:               pc.ID,
		Name:             pc.Name,
		Symbol:           pc.Symbol,
		Slug:             pc.Slug,
		LastUpdated:      parseTime(pc.LastUpdated),
		DateAdded:        parseTime(pc.DateAdded),
		DateLaunched:     parseTime(pc.DateLaunched),
		Price:            decimal.NewFromFloat(pc.Quote.USD.Price),
		Volume24h:        decimal.NewFromFloat(pc.Quote.USD.Volume24h),
		VolumeChange24h:  decimal.NewFromFloat(pc.Quote.USD.VolumeChange24h),
		PercentChange1h:  decimal.NewFromFloat(pc.Quote.USD.PercentChange1h),
		PercentChange24h: decimal.NewFromFloat(pc.Quote.USD.PercentChange24h),
		PercentChange7d:  decimal.NewFromFloat(pc.Quote.USD.PercentChange7d),
		MarketCap:        decimal.NewFromFloat(pc.Quote.USD.MarketCap),
		MarketCapDom:     decimal.NewFromFloat(pc.Quote.USD.MarketCapDom),
		FullyDilutedMC:   decimal.NewFromFloat(pc.Quote.USD.FullyDilutedMC),
		DBurls: models.DBurls{
			Website:      joinURLs(pc.Urls.Website),
			TechnicalDoc: joinURLs(pc.Urls.TechnicalDoc),
			Twitter:      joinURLs(pc.Urls.Twitter),
			Reddit:       joinURLs(pc.Urls.Reddit),
			MessageBoard: joinURLs(pc.Urls.MessageBoard),
			Announcement: joinURLs(pc.Urls.Announcement),
			Chat:         joinURLs(pc.Urls.Chat),
			Explorer:     joinURLs(pc.Urls.Explorer),
			SourceCode:   joinURLs(pc.Urls.SourceCode),
		},
		Logo:        pc.Logo,
		Description: pc.Description,
	}
}
