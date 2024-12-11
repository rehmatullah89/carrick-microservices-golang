package db

import (
	"carrick-js-api/pkgs/config"
	"carrick-js-api/pkgs/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type dbSingleton struct {
	dbConnection *gorm.DB
}

var (
	dbOnce     sync.Once
	dbInstance *dbSingleton
)

func GetDBInstance() *dbSingleton {
	if dbInstance == nil {
		dbOnce.Do(
			func() {
				dbInstance = &dbSingleton{}
				dbInstance.connect()
			})
	} else {
		dbInstance.ReConnect()
	}

	return dbInstance
}

func (dbI *dbSingleton) ReConnect() *dbSingleton {
	if dbI.dbConnection == nil {
		dbI.connect()
		return dbI
	}

	if db, err := dbI.dbConnection.DB(); err != nil {
		dbI.connect()
		return dbI
	} else {
		if err := db.Ping(); err != nil {
			dbI.connect()
			return dbI
		}
	}

	return dbI
}

func (dbI *dbSingleton) connect() {
	logger := logger.GetLoggerInstance()

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s TimeZone=UTC",
		config.AppConfig.DB.Host, config.AppConfig.DB.User,
		config.AppConfig.DB.Name, config.AppConfig.DB.Port,
		config.AppConfig.DB.Password)

	connect, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("DB connected")
		dbI.dbConnection = connect
	}
}

func (dbI *dbSingleton) GetDB() *gorm.DB {
	return dbI.dbConnection
}