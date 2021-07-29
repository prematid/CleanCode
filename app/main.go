package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"github.com/prematid/CleanCode/config"
	"github.com/prematid/CleanCode/stickers/delivery/http"
	"github.com/prematid/CleanCode/stickers/repository/mysql"
	"github.com/prematid/CleanCode/stickers/usecase"
)

var DB *gorm.DB

//main function
func main() {
	// create a new echo instance
	e := echo.New()
	var err error
	config.BuilDBConfig()

	DB, err = gorm.Open(config.DbConfig.DBtype, config.DbConfig.DBURL)

	if err != nil {
		log.Panic(err)
	}
	defer DB.Close()
	fmt.Println("Database Connection Successful")

	st := mysql.NewMysqlStickerRepository(DB)

	su := usecase.NewStickerUsecase(st)
	http.NewStickerHandler(e, su)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
