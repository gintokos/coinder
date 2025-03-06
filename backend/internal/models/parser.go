package models

type ParserCoin struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
	Slug        string      `json:"slug"`
	CMCRank     int         `json:"cmc_rank"`
	NumPairs    int         `json:"num_market_pairs"`
	CircSupply  float64     `json:"circulating_supply"`
	TotalSupply float64     `json:"total_supply"`
	MaxSupply   float64     `json:"max_supply"`
	Infinite    bool        `json:"infinite_supply"`
	LastUpdated string      `json:"last_updated"`
	DateAdded   string      `json:"date_added"`
	Tags        []string    `json:"tags"`
	SelfRepCS   *float64    `json:"self_reported_circulating_supply"`
	SelfRepMC   *float64    `json:"self_reported_market_cap"`
	Quote       ParserQuote `json:"quote"`

	ParserMetaDataCoin
}

type ParserQuote struct {
	USD ParserUsd `json:"USD"`
}

type ParserUsd struct {
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume_24h"`
	VolumeChange24h  float64 `json:"volume_change_24h"`
	PercentChange1h  float64 `json:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	PercentChange7d  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	MarketCapDom     float64 `json:"market_cap_dominance"`
	FullyDilutedMC   float64 `json:"fully_diluted_market_cap"`
	LastUpdated      string  `json:"last_updated"`
}

type ParserMetaDataCoin struct {
	Urls         ParserURLs `json:"urls"`
	Logo         string     `json:"logo"`
	Description  string     `json:"description"`
	Notice       *string    `json:"notice"`
	DateLaunched string     `json:"date_launched"`
	Category     string     `json:"category"`
	SelfRepTags  *[]string  `json:"self_reported_tags"`
}

type ParserURLs struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Twitter      []string `json:"twitter"`
	Reddit       []string `json:"reddit"`
	MessageBoard []string `json:"message_board"`
	Announcement []string `json:"announcement"`
	Chat         []string `json:"chat"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
}
