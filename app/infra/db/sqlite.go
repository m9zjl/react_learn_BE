package db

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	log "log/slog"
	"time"

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Error("fail to connect to sqlite err:%v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("failed to open db:%v", err)
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	// ping
	err = sqlDB.Ping()
	if err != nil {
		log.Error("failed to ping db:%v", err)
	} else {
		log.Info("connected to db")
	}
}
