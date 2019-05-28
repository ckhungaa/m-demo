package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DatabasePassword = "P@ssW0rd"
	DataBaseHost     = "m-mysql.demo"
	DatabasePort     = 3306
	DatabaseSchema   = "demo"
)

type Student struct {
	gorm.Model
	Name  string
	Class string
}

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("root:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DatabasePassword, DataBaseHost, DatabasePort, DatabaseSchema))
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&Student{}).Error; err != nil {
		panic("failed to migrate database")
	}
	return db
}
