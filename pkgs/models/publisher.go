package models

import (
	"carrick-js-api/pkgs/cache"
	"carrick-js-api/pkgs/config"
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/helpers"
	"errors"
	"regexp"
)

type Publisher struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"size:255;unique;not null"`
	Hash   string `gorm:"size:32;unique;not null"`
	Domain string `gorm:"size:100;not null;default:*"`
	Tag    string `gorm:"size:50"`
	db.Timestamps

	Tags          []Tag             `gorm:"constraint:OnDelete:CASCADE;"`
	Urls          []Url             `gorm:"constraint:OnDelete:CASCADE;"`
	PublisherUrls []PublisherUrl    `gorm:"constraint:OnDelete:CASCADE;"`
	Visits        []Visit           `gorm:"constraint:OnDelete:CASCADE;"`
	Domains       []PublisherDomain `gorm:"constraint:OnDelete:CASCADE;"`
}

func PublisherHasUrl(publisherHash string, rawUrl string) bool {
	db := db.GetDBInstance().GetDB()
	cache := cache.GetRedisCacheInstance()

	var count int64

	urlPath, err := helpers.GetPathFromUrl(rawUrl)
	if err != nil {
		return false
	}

	cacheKey := publisherHash + urlPath
	if err := cache.Get(cacheKey, &count); err == nil && count > 0 {
		return true
	}

	sql := `select count(1)
		from urls u
		left join publishers p on p.id = u.publisher_id
		where p.hash = @publisher_hash
			and u.url_path = @url_path`
	sqlParams := map[string]interface{}{
		"publisher_hash": publisherHash,
		"url_path":       urlPath,
	}
	db.Raw(sql, sqlParams).Count(&count)
	if count == 0 {
		return false
	}

	cache.Set(cacheKey, count, config.AppConfig.CacheTTL)
	return true
}

func PublisherByHash(publisherHash string) (Publisher, error) {
	db := db.GetDBInstance().GetDB()
	cache := cache.GetRedisCacheInstance()

	var publisher Publisher
	if err := cache.Get(publisherHash, &publisher); err == nil && publisher.ID != 0 {
		return publisher, nil
	}

	db.Model(Publisher{}).Preload("Domains").Where("hash = ?", publisherHash).
		Where("deleted_at IS NULL").
		First(&publisher)

	if publisher.ID != 0 {
		cache.Set(publisherHash, publisher, config.AppConfig.CacheTTL)
		return publisher, nil
	} else {
		return Publisher{}, errors.New("Publisher not found by hash: " + publisherHash)
	}
}

func PublisherHasDomain(domain string, publisher Publisher) bool {
	for _, d := range publisher.Domains {
		domainMatch, err := regexp.MatchString(`.*`+d.Domain+`$`, domain)
		if err == nil && domainMatch {
			return true
		}
	}

	return false
}
