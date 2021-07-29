package http

import (
	"fmt"
	"net/http"
	"strconv"

	//"strconv"

	"github.com/labstack/echo"
	"github.com/prematid/CleanCode/domain"
)

type StickerHandler struct {
	StickerUsecase domain.StickerUsecase
}

func NewStickerHandler(e *echo.Echo, us domain.StickerUsecase) {
	handler := &StickerHandler{StickerUsecase: us}
	e.GET("/", Welcome)

	e.GET("/v1/trendingStickers", handler.GetSticker)
}

func Welcome(c echo.Context) error {
	fmt.Println("Welcome Method")
	return c.String(http.StatusOK, "Hello, World! this is test")
}

func (s *StickerHandler) GetSticker(c echo.Context) error {
	l := 10
	p := 1

	var err error

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	if limit != "" {
		l, err = strconv.Atoi(limit)
	}
	if page != "" {
		p, err = strconv.Atoi(page)
	}

	fmt.Println("Trending Stickers Method")
	stickerss, _ := s.StickerUsecase.GetStickers(l, p)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, stickerss)
}
