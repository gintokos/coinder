package models

import (
	"github.com/shopspring/decimal"
)

type DBCoin struct {
	ID     int    `gorm:"primaryKey;type:integer" json:"id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Symbol string `gorm:"type:varchar(20);not null" json:"symbol"`
	Slug   string `gorm:"type:varchar(100);" json:"slug"`

	LastUpdated  *string `gorm:"type:timestamp" json:"last_updated"`
	DateAdded    *string `gorm:"type:timestamp" json:"date_added"`
	DateLaunched *string `gorm:"type:timestamp" json:"date_launched"`

	Price            decimal.Decimal `gorm:"type:decimal" json:"price"`
	Volume24h        decimal.Decimal `gorm:"type:decimal" json:"volume_24h"`
	VolumeChange24h  decimal.Decimal `gorm:"type:decimal" json:"volume_change_24h"`
	PercentChange1h  decimal.Decimal `gorm:"type:decimal" json:"percent_change_1h"`
	PercentChange24h decimal.Decimal `gorm:"type:decimal" json:"percent_change_24h"`
	PercentChange7d  decimal.Decimal `gorm:"type:decimal" json:"percent_change_7d"`
	MarketCap        decimal.Decimal `gorm:"type:decimal" json:"market_cap"`
	MarketCapDom     decimal.Decimal `gorm:"type:decimal" json:"market_cap_dominance"`
	FullyDilutedMC   decimal.Decimal `gorm:"type:decimal" json:"fully_diluted_market_cap"`

	DBurls      DBurls `gorm:"embedded" json:"urls"`
	Logo        string `gorm:"type:text" json:"logo"`
	Description string `gorm:"type:text" json:"description"`
}

type DBurls struct {
	Website      string `gorm:"type:text" json:"website"`
	TechnicalDoc string `gorm:"type:text" json:"technical_doc"`
	Twitter      string `gorm:"type:text" json:"twitter"`
	Reddit       string `gorm:"type:text" json:"reddit"`
	MessageBoard string `gorm:"type:text" json:"message_board"`
	Announcement string `gorm:"type:text" json:"announcement"`
	Chat         string `gorm:"type:text" json:"chat"`
	Explorer     string `gorm:"type:text" json:"explorer"`
	SourceCode   string `gorm:"type:text" json:"source_code"`
}


