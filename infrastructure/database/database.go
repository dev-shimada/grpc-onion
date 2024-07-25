package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// dbコンテナが立ち上がるまで接続を行う
func RetryConnectDB(dialector gorm.Dialector, opt gorm.Option, count uint) error {
	var err error
	for count > 1 {
		if db, err = gorm.Open(dialector, opt); err != nil {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... coutn:%v\n", count)
			continue
		}
		break
	}
	return err
}

func NewDB() *gorm.DB {
	// sqliteに接続
	if err := RetryConnectDB(sqlite.Open("onion.db"), &gorm.Config{}, 100); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
