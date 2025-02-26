package models

import (
	"github.com/shopspring/decimal"
)

type SearchCoinOpt struct {
	SearchCoinOptReq
	UserIDLClient int64
}

type SearchCoinOptReq struct {
	UserIDTarget int64  `json:"user_id_target,omitempty" validate:"omitempty,gt=0"`
	Page         int    `json:"page" validate:"required,min=1,max=1000"`
	Limit        int    `json:"limit" validate:"required,min=1,max=100"`
	LikedByUser  bool   `json:"liked_by_user"`
	SortedBy     string `json:"sorted_by" validate:"required,oneof=BY_PRICE BY_MARKET_CAP BY_POPULARITY"`
	LikedToday   bool   `json:"liked_today"`
}

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

	LikesCount    int `gorm:"type:integer;default:0" json:"likes_count"`
	CommentsCount int `gorm:"type:integer;default:0" json:"comments_count"`
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

func (DBCoin) TableName() string {
	return "coins"
}
