package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	PhotoUrl  string    `json:"photo_url"`
	AuthDate  string    `json:"auth_date"`
	CreatedAt time.Time `json:"created_at"`
}

type coinuser struct {
	CoinID int   `json:"coin_id"`
	UserID int64 `json:"user_id"`
}

type Likes struct {
	coinuser `gorm:"embedded;embeddedPrefix:like_"`
}

type Favorite struct {
	coinuser `gorm:"embedded;embeddedPrefix:favorite_"`
}

type Share struct {
	coinuser `gorm:"embedded;embeddedPrefix:share_"`
}

type Comment struct {
	coinuser `gorm:"embedded;embeddedPrefix:comment_"`
}
