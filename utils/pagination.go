package utils

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/prematid/CleanCode/domain"

)

//GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c echo.Context) domain.Pagination {
	// Initializing default
	//	var mode string
	l :=2
	p := 1
	sort := "created_at asc"
	limit:= c.QueryParam("limit")
	page := c.QueryParam("page")

	l, err := strconv.Atoi(limit)

	if err != nil {
		panic(err)
	}
   
	p, err = strconv.Atoi(page)

	if err != nil {
		panic(err)
	}
	

	return domain.Pagination{
		Limit: l,
		Page:  p,
		Sort:  sort,
	}

}
