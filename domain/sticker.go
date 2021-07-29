package domain

import (
	"time"
)

type Stickerss struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	Priority  int       `json:"priority"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type SendSticker struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	Priority  int       `json:"priority"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

//sticker usecase interface
type StickerUsecase interface {
	GetStickers(limit int, page int) ([]Stickerss, error)
}

//sticker repository contract interface
type StickerRepository interface {
	GetStickers(limit int, page int) ([]Stickerss, error)
}
type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
