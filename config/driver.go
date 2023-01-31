package config

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
  "crud-echo/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbConfig := Config.db
	fmt.Printf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	if dbConfig.Adapter == "postgres" {
		db, err = gorm.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Name))
		log.Println("Connected to Database Local postgressql")
	}

	if err != nil {
		log.Println("[Driver.ConnectDB] error when connect to database")
		log.Fatal("[Driver.ConnectDB] error when connect to database")
	} else {
		log.Println("SUCCESS CONNECT TO DATABASE")
	}

	go doEvery(6*time.Minute, pingDb, db)

	// Database Pooling
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)
	db.DB().SetConnMaxLifetime(45 * time.Second)

  models.InitTableCustomer(db)

	return db
}

func doEvery(d time.Duration, f func(*gorm.DB), y *gorm.DB) {
	for _ = range time.Tick(d) {
		f(y)
	}
}

func pingDb(db *gorm.DB) {
	log.Println("PING CONNECTION")
	err := db.DB().Ping()
	if err != nil {
		log.Println("PING CONNECTION FAILURE")
		return
	}
}