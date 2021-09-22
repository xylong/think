package db

import (
	"log"
	"think/gin/src/model/LogModel"
	"think/gin/src/model/UserModel"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Orm *gorm.DB

func InitDB() {
	Orm = gormDB()
}

func InitTable() {
	if err := Orm.AutoMigrate(
		UserModel.New(),
		LogModel.New(),
	); err != nil {
		log.Fatalln(err)
	}
}

func gormDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetConnMaxLifetime(time.Second * 30)

	return db
}
