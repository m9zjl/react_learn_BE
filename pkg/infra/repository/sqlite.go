package repository

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	log "log/slog"
	"os"
	"server/pkg/domain/entity"
	"time"
)

func InitDB() *gorm.DB {
	db, err := connectDatabase()
	if err != nil {
		log.Error("failed to connect to db")
		os.Exit(-1)
	}
	return db
}

func connectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./assets/db/db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Error("fail to connect to sqlite err:%v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("failed to open DB:%v", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	// ping
	err = sqlDB.Ping()
	if err != nil {
		log.Error("failed to ping DB:%v", err)
		return nil, err
	} else {
		log.Info("connected to DB")
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Error("failed to migrate user")
	}

	err = db.AutoMigrate(&entity.Article{})
	if err != nil {
		log.Error("failed to migrate article")
	}
	return db, err
}
