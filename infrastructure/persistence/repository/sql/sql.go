package sql

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func StartUp(url string) {
	var err error
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(20)
	DB.SingularTable(true)
}
