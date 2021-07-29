package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/prematid/CleanCode/domain"
)

type mysqlStickerRepository struct {
	Conn *gorm.DB
}

func NewMysqlStickerRepository(Conn *gorm.DB) domain.StickerRepository {
	return &mysqlStickerRepository{Conn}
}

func (s *mysqlStickerRepository) GetStickers(limit int, page int) ([]domain.Stickerss, error) {

	stickers := []domain.Stickerss{}
	offset := (page - 1) * limit

	fmt.Println(offset)
	fmt.Println(stickers)
	if err := s.Conn.Limit(limit).Offset(offset).Order("priority, id").Find(&stickers).Error; err != nil {
		return nil, err
	}

	fmt.Println("Yes after")
	return stickers, nil
}
