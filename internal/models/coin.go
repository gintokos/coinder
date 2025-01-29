package models

type DBCoin struct {
	ID     string `gorm:"primaryKey;type:varchar(100)" json:"id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Symbol string `gorm:"type:varchar(20);not null" json:"symbol"`
	Slug   string `gorm:"type:varchar(100);uniqueIndex" json:"slug"`

	LastUpdated string `gorm:"type:timestamp" json:"last_updated"`
	DateAdded   string `gorm:"type:timestamp" json:"date_added"`

	Price            float64 `gorm:"type:decimal(20,8)" json:"price"`
	Volume24h        float64 `gorm:"type:decimal(20,2)" json:"volume_24h"`
	VolumeChange24h  float64 `gorm:"type:decimal(20,2)" json:"volume_change_24h"`
	PercentChange1h  float64 `gorm:"type:decimal(10,2)" json:"percent_change_1h"`
	PercentChange24h float64 `gorm:"type:decimal(10,2)" json:"percent_change_24h"`
	PercentChange7d  float64 `gorm:"type:decimal(10,2)" json:"percent_change_7d"`
	MarketCap        float64 `gorm:"type:decimal(20,2)" json:"market_cap"`
	MarketCapDom     float64 `gorm:"type:decimal(10,2)" json:"market_cap_dominance"`
	FullyDilutedMC   float64 `gorm:"type:decimal(20,2)" json:"fully_diluted_market_cap"`

	DBurls       DBurls `gorm:"embedded" json:"urls"`
	Logo         string `gorm:"type:text" json:"logo"`
	Description  string `gorm:"type:text" json:"description"`
	DateLaunched string `gorm:"type:timestamp" json:"date_launched"`
}

type DBurls struct {
	Website      []string `gorm:"type:text[]" json:"website"`
	TechnicalDoc []string `gorm:"type:text[]" json:"technical_doc"`
	Twitter      []string `gorm:"type:text[]" json:"twitter"`
	Reddit       []string `gorm:"type:text[]" json:"reddit"`
	MessageBoard []string `gorm:"type:text[]" json:"message_board"`
	Announcement []string `gorm:"type:text[]" json:"announcement"`
	Chat         []string `gorm:"type:text[]" json:"chat"`
	Explorer     []string `gorm:"type:text[]" json:"explorer"`
	SourceCode   []string `gorm:"type:text[]" json:"source_code"`
}
