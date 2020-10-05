package db

import (
	"fmt"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

//DB is
var DB *gorm.DB

//Init is used to Initialize Database
func Init() (*gorm.DB, error) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("podgrab.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}

	localDB, _ := db.DB()
	localDB.SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB, nil
}

//Migrate Database
func Migrate() {
	DB.AutoMigrate(&Podcast{}, &PodcastItem{})

}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
