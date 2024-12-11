package models

import (
	"carrick-js-api/pkgs/db"
	"database/sql"
	"errors"
)

type TrafficSource struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:50;not null;unique"`
	Clid sql.NullString `gorm:"size:10"`
	Is_Default bool `gorm:"default:false"`
	Is_Amp bool `gorm:"default:false"`
	db.Timestamps

	Domains []TrafficSourceDomain `gorm:"constraint:OnDelete:CASCADE;"`
	PublisherUrls []PublisherUrl `gorm:"constraint:OnDelete:CASCADE;"`
	Visits []Visit `gorm:"constraint:OnDelete:CASCADE;"`
}

func GetTrafficSourceByDomain(domain string) (TrafficSource, error) {
	db := db.GetDBInstance().GetDB()

	var trafficSourceResult TrafficSource
	db.Raw("select * from traffic_source_by_domain(?)", domain).First(&trafficSourceResult)

	if trafficSourceResult.ID == 0 {
		return TrafficSource{}, errors.New("Traffic source not found")
	}

	return trafficSourceResult, nil
}