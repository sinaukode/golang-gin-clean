package config

import (
	"github.com/jinzhu/gorm"
	"go-clean-code-gin/entity"
	"log"
)

func DbConnect() *gorm.DB {
	str := "root:@tcp(127.0.0.1:3306)/simple_api?parseTime=true"
	db, err := gorm.Open("mysql", str)

	if err != nil {
		log.Fatal("Error when connect db" + str + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(
		entity.Person{},
	)

	return db

}
