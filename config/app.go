package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// create variable to connect with database
func Connect() {
	d, err := gorm.Open("mysql", "root:india@123@/studentinfor?charset=utf8&parseTime=True&loc=Local") //ORM library for dealing with relational databases

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
