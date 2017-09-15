package db

import (
	"github.com/jinzhu/gorm"
	"../config"
)


func Init()  {

	db,err := gorm.Open("mysql",config.DB_USERNAME + ":" + config.DB_PASSWORD + "@" + config.DB_HOST + ":" + config.DB_PASSWORD + "/" + config.DB_NAME)

	if err == nil {
		panic(err.Error())
	}
	gorm.
}
