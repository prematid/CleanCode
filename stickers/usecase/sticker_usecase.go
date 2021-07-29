package usecase

import (
	"fmt"

	"github.com/prematid/CleanCode/domain"
)

type stickerUsecase struct {
	StickerRepo domain.StickerRepository
}

func NewStickerUsecase(a domain.StickerRepository) domain.StickerUsecase {
	return &stickerUsecase{
		StickerRepo: a,
	}
}

func (s *stickerUsecase) GetStickers(limit int, page int) (stickers []domain.Stickerss, err error) {
	stickers = []domain.Stickerss{}

	fmt.Println(stickers)
	offset := (page - 1) * limit
	fmt.Println(offset)

	stickers, err = s.StickerRepo.GetStickers(limit, page)
	fmt.Println(stickers)

	return stickers, err
}
