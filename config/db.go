package config

import "fmt"

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     int
	DBtype     string
	DBURL      string
}

var DbConfig DBConfig

func BuilDBConfig() {
	DbConfig = DBConfig{
		DBUser:     "root",
		DBPassword: "root",
		DBName:     "sticker",
		DBHost:     "localhost",
		DBPort:     3306,
		DBtype:     "mysql",
	}
	DbConfig.DBURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbConfig.DBUser,
		DbConfig.DBPassword,
		DbConfig.DBHost,
		DbConfig.DBPort,
		DbConfig.DBName)
}
