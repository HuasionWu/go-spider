package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB       *gorm.DB
	username string = "root"
	password string = "123456"
	dbName   string = "spiders"
)

func init() {
	var err error

	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))

	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sp_" + defaultTableName
	}

	DB.SingularTable(true)
}

func CloseDB() {
	defer DB.Close()
}
