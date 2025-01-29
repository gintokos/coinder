package models

import "strconv"

func ToDBcoin(pc *ParserCoin) DBCoin {
	return DBCoin{
		ID:          strconv.Itoa(pc.ID),
		Name:        pc.Name,
		Symbol:      pc.Symbol,
		Slug:        pc.Slug,
		LastUpdated: pc.LastUpdated,
		DateAdded:   pc.DateAdded,

		Price:            pc.Quote.USD.Price,
		Volume24h:        pc.Quote.USD.Volume24h,
		VolumeChange24h:  pc.Quote.USD.VolumeChange24h,
		PercentChange1h:  pc.Quote.USD.PercentChange1h,
		PercentChange24h: pc.Quote.USD.PercentChange24h,
		PercentChange7d:  pc.Quote.USD.PercentChange7d,
		MarketCap:        pc.Quote.USD.MarketCap,
		MarketCapDom:     pc.Quote.USD.MarketCapDom,
		FullyDilutedMC:   pc.Quote.USD.FullyDilutedMC,

		DBurls: DBurls{
			Website:      pc.Urls.Website,
			TechnicalDoc: pc.Urls.TechnicalDoc,
			Twitter:      pc.Urls.Twitter,
			Reddit:       pc.Urls.Reddit,
			MessageBoard: pc.Urls.MessageBoard,
			Announcement: pc.Urls.Announcement,
			Chat:         pc.Urls.Chat,
			Explorer:     pc.Urls.Explorer,
			SourceCode:   pc.Urls.SourceCode,
		},
		Logo:         pc.Logo,
		Description:  pc.Description,
		DateLaunched: pc.DateLaunched,
	}
}
