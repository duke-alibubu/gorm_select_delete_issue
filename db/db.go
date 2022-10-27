package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB
var dsnFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
var dsn = fmt.Sprintf(dsnFormat, "root", "root", "127.0.0.1", "3306", "playground")

func MustInit() {
	var err error
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbConn = dbConn.Debug()
}

func GetDB() *gorm.DB {
	return dbConn
}
