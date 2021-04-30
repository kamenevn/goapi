package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kamenevn/goapi/helpers"
)

var DB *gorm.DB

func InitDb(dbType string, Config *Config) *gorm.DB {
	//if INITDBPG == false {
		//PrintToConsole("Postgres грузим")
		//DB = SetupDb("postgres", CONFIG, "balance")

	var connectString string
		connectString = fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v sslmode=disable",
			Config.Database.Host,
			Config.Database.User,
			Config.Database.Password,
			Config.Database.Name,
		)

	db, err := gorm.Open(dbType, connectString)
	helpers.CheckErr(err, "Ошибка подключения к " + dbType)

	db.SingularTable(false)

	//defer db.Close()
	db.LogMode(true)
	//TX = db.Begin()

	return db
}