package main

import (
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/models"
	"gorm.io/gorm"
)

func publisherTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.Publisher{})
}

func tagTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.Tag{})
}

func urlTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.Url{})
}

func trafficSourceTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.TrafficSource{})

	db.Migrator().CreateTable(&models.TrafficSourceDomain{})
}

func publisherUrlTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.PublisherUrl{})
}

func visitTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.Visit{})
}

func clicksTrackingBufferTable(db *gorm.DB) {
	db.Migrator().CreateTable(&models.ClicksTrackingBuffer{})
}

func main() {
	db := db.GetDBInstance().GetDB()

	publisherTable(db)
	tagTable(db)
	urlTable(db)
	trafficSourceTable(db)
	publisherUrlTable(db)
	visitTable(db)
	clicksTrackingBufferTable(db)
}