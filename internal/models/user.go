package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	PhotoUrl  string    `json:"photo_url"`
	AuthDate  int64     `json:"auth_date"`
	CreatedAt time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}

type Coinuser struct {
	CoinID    int       `json:"coin_id" gorm:"primaryKey;index:idx_coin_user,priority:1"`
	UserID    int64     `json:"user_id" gorm:"primaryKey;index:idx_coin_user,priority:2"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	User User   `gorm:"foreignKey:UserID;references:ID"`
	Coin DBCoin `gorm:"foreignKey:CoinID;references:ID"`
}

type Likes struct {
	Coinuser `gorm:"embedded;embeddedPrefix:like_"`
}

func (Likes) TableName() string {
	return "likes"
}

type Favorite struct {
	Coinuser `gorm:"embedded;embeddedPrefix:favorite_"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type Share struct {
	Coinuser `gorm:"embedded;embeddedPrefix:share_"`
}

func (Share) TableName() string {
	return "shares"
}

type Comment struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CoinID    int       `json:"coin_id" gorm:"index"`
	UserID    int64     `json:"user_id" gorm:"index"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	User User   `gorm:"foreignKey:UserID;references:ID"`
	Coin DBCoin `gorm:"foreignKey:CoinID;references:ID"`
}

func (Comment) TableName() string {
	return "comments"
}
