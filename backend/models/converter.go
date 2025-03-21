package models

import (
	"strings"

	pb "github.com/gintokos/coinder/protos/coinupdateprotos"
	"github.com/shopspring/decimal"
)

func ToPBCoin(coins []DBCoin) *pb.Coins {
	result := &pb.Coins{
		Coins: make([]*pb.Coin, len(coins)),
	}

	for i, coin := range coins {
		pbCoin := &pb.Coin{
			Id:          int32(coin.ID),
			Name:        coin.Name,
			Symbol:      coin.Symbol,
			Slug:        coin.Slug,
			LastUpdated: safeString(coin.LastUpdated),
			DateAdded:   safeString(coin.DateAdded),

			Usd: &pb.Usd{
				Price:            float32(coin.Price.InexactFloat64()),
				Volume24H:        float32(coin.Volume24h.InexactFloat64()),
				VolumeChange24H:  float32(coin.VolumeChange24h.InexactFloat64()),
				PercentChange1H:  float32(coin.PercentChange1h.InexactFloat64()),
				PercentChange24H: float32(coin.PercentChange24h.InexactFloat64()),
				PercentChange7D:  float32(coin.PercentChange7d.InexactFloat64()),
				MarketCap:        float32(coin.MarketCap.InexactFloat64()),
				MarketCapDom:     float32(coin.MarketCapDom.InexactFloat64()),
				FullyDilutedMC:   float32(coin.FullyDilutedMC.InexactFloat64()),
				LastUpdated:      safeString(coin.LastUpdated),
			},

			MetaData: &pb.MetaData{
				Logo:         coin.Logo,
				Description:  coin.Description,
				DateLaunched: safeString(coin.DateLaunched),
				Urls:         convertUrls(coin.DBurls),
			},
		}

		result.Coins[i] = pbCoin
	}

	return result
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func convertUrls(dbUrls DBurls) []*pb.Url {
	websiteUrl := &pb.Url{
		Urls: []string{dbUrls.Website},
	}

	technicalDocUrl := &pb.Url{
		Urls: []string{dbUrls.TechnicalDoc},
	}

	twitterUrl := &pb.Url{
		Urls: []string{dbUrls.Twitter},
	}

	redditUrl := &pb.Url{
		Urls: []string{dbUrls.Reddit},
	}

	messageBoardUrl := &pb.Url{
		Urls: []string{dbUrls.MessageBoard},
	}

	announcementUrl := &pb.Url{
		Urls: []string{dbUrls.Announcement},
	}

	chatUrl := &pb.Url{
		Urls: []string{dbUrls.Chat},
	}

	explorerUrl := &pb.Url{
		Urls: []string{dbUrls.Explorer},
	}

	sourceCodeUrl := &pb.Url{
		Urls: []string{dbUrls.SourceCode},
	}

	urls := []*pb.Url{
		websiteUrl,
		technicalDocUrl,
		twitterUrl,
		redditUrl,
		messageBoardUrl,
		announcementUrl,
		chatUrl,
		explorerUrl,
		sourceCodeUrl,
	}

	filteredUrls := make([]*pb.Url, 0)
	for _, url := range urls {
		if len(url.Urls) > 0 && url.Urls[0] != "" {
			filteredUrls = append(filteredUrls, url)
		}
	}

	return filteredUrls
}

func ToDBCoinsFromPB(pbCoins *pb.Coins) []DBCoin {
	if pbCoins == nil || len(pbCoins.Coins) == 0 {
		return []DBCoin{}
	}

	result := make([]DBCoin, len(pbCoins.Coins))

	for i, pbCoin := range pbCoins.Coins {
		dbCoin := DBCoin{
			CommonInfoCoin: CommonInfoCoin{
				ID:           int(pbCoin.Id),
				Name:         pbCoin.Name,
				Symbol:       pbCoin.Symbol,
				Slug:         pbCoin.Slug,
				LastUpdated:  strPtr(pbCoin.LastUpdated),
				DateAdded:    strPtr(pbCoin.DateAdded),
				DateLaunched: strPtr(extractDateLaunched(pbCoin.MetaData)),
				Logo:         extractLogo(pbCoin.MetaData),
				Description:  extractDescription(pbCoin.MetaData),
			},
			CoinStats: CoinStats{},
		}

		if pbCoin.Usd != nil {
			dbCoin.Price = decimal.NewFromFloat32(pbCoin.Usd.Price)
			dbCoin.Volume24h = decimal.NewFromFloat32(pbCoin.Usd.Volume24H)
			dbCoin.VolumeChange24h = decimal.NewFromFloat32(pbCoin.Usd.VolumeChange24H)
			dbCoin.PercentChange1h = decimal.NewFromFloat32(pbCoin.Usd.PercentChange1H)
			dbCoin.PercentChange24h = decimal.NewFromFloat32(pbCoin.Usd.PercentChange24H)
			dbCoin.PercentChange7d = decimal.NewFromFloat32(pbCoin.Usd.PercentChange7D)
			dbCoin.MarketCap = decimal.NewFromFloat32(pbCoin.Usd.MarketCap)
			dbCoin.MarketCapDom = decimal.NewFromFloat32(pbCoin.Usd.MarketCapDom)
			dbCoin.FullyDilutedMC = decimal.NewFromFloat32(pbCoin.Usd.FullyDilutedMC)
		}

		if pbCoin.MetaData != nil {
			dbCoin.DBurls = extractUrls(pbCoin.MetaData.Urls)
		}

		result[i] = dbCoin
	}

	return result
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func extractDateLaunched(metaData *pb.MetaData) string {
	if metaData == nil {
		return ""
	}
	return metaData.DateLaunched
}

func extractLogo(metaData *pb.MetaData) string {
	if metaData == nil {
		return ""
	}
	return metaData.Logo
}

func extractDescription(metaData *pb.MetaData) string {
	if metaData == nil {
		return ""
	}
	return metaData.Description
}

func extractUrls(pbUrls []*pb.Url) DBurls {
	result := DBurls{}

	urlMap := make(map[string]string)

	for _, urlObj := range pbUrls {
		for _, url := range urlObj.Urls {

			if strings.Contains(url, "twitter.com") {
				urlMap["twitter"] = url
			} else if strings.Contains(url, "reddit.com") {
				urlMap["reddit"] = url
			} else if strings.Contains(url, "github.com") {
				urlMap["source_code"] = url
			} else if strings.Contains(url, ".pdf") || strings.Contains(url, "whitepaper") {
				urlMap["technical_doc"] = url
			} else if strings.Contains(url, "explorer") || strings.Contains(url, "scan") {
				urlMap["explorer"] = url
			} else if strings.Contains(url, "t.me") || strings.Contains(url, "discord") {
				urlMap["chat"] = url
			} else if strings.Contains(url, "medium.com") || strings.Contains(url, "blog") {
				urlMap["message_board"] = url
			} else if strings.Contains(url, "announce") {
				urlMap["announcement"] = url
			} else {
				if urlMap["website"] == "" {
					urlMap["website"] = url
				}
			}
		}
	}
	result.Website = urlMap["website"]
	result.TechnicalDoc = urlMap["technical_doc"]
	result.Twitter = urlMap["twitter"]
	result.Reddit = urlMap["reddit"]
	result.MessageBoard = urlMap["message_board"]
	result.Announcement = urlMap["announcement"]
	result.Chat = urlMap["chat"]
	result.Explorer = urlMap["explorer"]
	result.SourceCode = urlMap["source_code"]

	return result
}
