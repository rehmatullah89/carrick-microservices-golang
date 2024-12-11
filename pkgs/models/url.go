package models

import (
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/helpers"
)

type Url struct {
	ID           uint   `gorm:"primarykey"`
	Publisher_Id uint   `gorm:"not null;index:urls_publisher_id_url_path_uidx,unique"`
	Url_Path     string `gorm:"size:600;not null;index:urls_publisher_id_url_path_uidx,unique"`
	db.Timestamps

	Publisher     Publisher      `gorm:"constraint:OnDelete:CASCADE;"`
	PublisherUrls []PublisherUrl `gorm:"constraint:OnDelete:CASCADE;"`
}

func FirstOrCreateUrl(publisher Publisher, rawUrl string) (Url, error) {
	db := db.GetDBInstance()

	urlPath, err := helpers.GetPathFromUrl(rawUrl)
	if err != nil {
		return Url{}, err
	}

	url := Url{
		Publisher_Id: publisher.ID,
		Url_Path:     urlPath,
	}

	if err := db.GetDB().FirstOrCreate(&url, Url{Publisher_Id: publisher.ID, Url_Path: urlPath}).Error; err != nil {
		return Url{}, err
	}

	return url, nil
}
