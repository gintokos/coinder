package parser

import (
	"strings"
	"time"

	"github.com/gintokos/coinder/backend/models"
	psModels "github.com/gintokos/coinder/coinparser/models"
	pb "github.com/gintokos/coinder/protos/coinupdateprotos"
	"github.com/shopspring/decimal"
)

func ToPBcoin(pc psModels.ParserCoin) *pb.Coin {
	usd := &pb.Usd{
		Price:            float32(pc.Quote.USD.Price),
		Volume24H:        float32(pc.Quote.USD.Volume24h),
		VolumeChange24H:  float32(pc.Quote.USD.VolumeChange24h),
		PercentChange1H:  float32(pc.Quote.USD.PercentChange1h),
		PercentChange24H: float32(pc.Quote.USD.PercentChange24h),
		PercentChange7D:  float32(pc.Quote.USD.PercentChange7d),
		MarketCap:        float32(pc.Quote.USD.MarketCap),
		MarketCapDom:     float32(pc.Quote.USD.MarketCapDom),
		FullyDilutedMC:   float32(pc.Quote.USD.FullyDilutedMC),
		LastUpdated:      pc.Quote.USD.LastUpdated,
	}

	var urls []*pb.Url

	addUrlsIfNotEmpty := func(urlSlice []string) {
		if len(urlSlice) > 0 {
			urls = append(urls, &pb.Url{
				Urls: urlSlice,
			})
		}
	}

	addUrlsIfNotEmpty(pc.Urls.Website)
	addUrlsIfNotEmpty(pc.Urls.TechnicalDoc)
	addUrlsIfNotEmpty(pc.Urls.Twitter)
	addUrlsIfNotEmpty(pc.Urls.Reddit)
	addUrlsIfNotEmpty(pc.Urls.MessageBoard)
	addUrlsIfNotEmpty(pc.Urls.Announcement)
	addUrlsIfNotEmpty(pc.Urls.Chat)
	addUrlsIfNotEmpty(pc.Urls.Explorer)
	addUrlsIfNotEmpty(pc.Urls.SourceCode)

	metaData := &pb.MetaData{
		Logo:         pc.Logo,
		Description:  pc.Description,
		Notice:       *pc.Notice,
		DateLaunched: pc.DateLaunched,
		Category:     pc.Category,
		Urls:         urls,
		SelfRepTags:  pc.Tags,
	}

	return &pb.Coin{
		Id:          int32(pc.ID),
		Name:        pc.Name,
		Symbol:      pc.Symbol,
		Slug:        pc.Slug,
		CircSupply:  float32(pc.CircSupply),
		TotalSupply: float32(pc.TotalSupply),
		MaxSupply:   float32(pc.MaxSupply),
		Infinite:    pc.Infinite,
		LastUpdated: pc.LastUpdated,
		DateAdded:   pc.DateAdded,
		Tags:        pc.Tags,
		SelfRepCS:   float32(*pc.SelfRepCS),
		SelfRepMC:   float32(*pc.SelfRepMC),
		Usd:         usd,
		MetaData:    metaData,
	}
}

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
		CoinStats: models.CoinStats{
			LikesCount:    0,
			CommentsCount: 0,
		},
		CommonInfoCoin: models.CommonInfoCoin{
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
		},
	}
}
