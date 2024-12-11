package models

import (
	"carrick-js-api/pkgs/cache"
	"carrick-js-api/pkgs/config"
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/logger"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type NewPublisherUrl struct {
	ID                uint          `gorm:"primarykey"`
	Publisher_Id      sql.NullInt64 `gorm:"default:null"`
	Url_Id            sql.NullInt64 `gorm:"default:null"`
	TagStr            string        `gorm:"column:tag;size:50;not null"`
	Traffic_Source_Id sql.NullInt64 `gorm:"default:null"`
	Created_At        time.Time     `gorm:"default:current_timestamp"`

	// TODO : No need for the constraints
	Publisher     Publisher     `gorm:"constraint:OnDelete:CASCADE"`
	Url           Url           `gorm:"constraint:OnDelete:CASCADE"`
	TrafficSource TrafficSource `gorm:"constraint:OnDelete:CASCADE"`
}

func (NewPublisherUrl) TableName() string {
	return "new_publisher_url"
}

func urlHasTag(publisher Publisher, urlPath string, trafficSource TrafficSource) (TagResult, error) {
	db := db.GetDBInstance().GetDB()
	cache := cache.GetRedisCacheInstance()

	var tagResult TagResult

	newPublisherUrlCacheKey := fmt.Sprintf("%d-%s-%d", publisher.ID, urlPath, trafficSource.ID)

	if err := cache.Get(newPublisherUrlCacheKey, &tagResult); err == nil && tagResult.Id != 0 {
		return tagResult, nil
	}

	sqlQuery := `SELECT pu.tag
		FROM new_publisher_url pu
		left join urls u on pu.url_id = u.id
		where pu.publisher_id=@publisher_id
			and pu.traffic_source_id=@traffic_source_id
			and u.url_path=@publisher_url
		`
	sqlParams := map[string]interface{}{
		"publisher_id":      publisher.ID,
		"publisher_url":     urlPath,
		"traffic_source_id": trafficSource.ID,
	}

	if err := db.Raw(sqlQuery, sqlParams).First(&tagResult).Error; err != nil {
		return TagResult{}, err
	}

	if tagResult.Id == 0 {
		return TagResult{}, errors.New("Tag for url not found")
	}

	cache.Set(newPublisherUrlCacheKey, tagResult, config.AppConfig.CacheTTL)

	return tagResult, nil
}

func AttachTagToUrlAndTrafficSource(publisher Publisher, trafficSource TrafficSource, urlPath string) (TagResult, error) {
	db := db.GetDBInstance().GetDB()
	logger := logger.GetLoggerInstance()
	var err error

	// create or find url
	url, err := FirstOrCreateUrl(publisher, urlPath)
	if err != nil || url.ID == 0 {
		logger.Error(err)
		return TagResult{}, err
	}

	if existsTagResult, _ := urlHasTag(publisher, urlPath, trafficSource); existsTagResult.Id != 0 {
		return existsTagResult, nil
	}

	// get unused tag
	tagResult := GetTagForTrafficType(publisher, "organic")

	if tagResult.Id == 0 {
		logger.Error("Free tag not found")
		return TagResult{}, errors.New("Free tag not found")
	}
	newPublisherUrl := NewPublisherUrl{
		Publisher_Id:      sql.NullInt64{Int64: int64(publisher.ID), Valid: publisher.ID != 0},
		Url_Id:            sql.NullInt64{Int64: int64(url.ID), Valid: url.ID != 0},
		TagStr:            tagResult.Tag,
		Traffic_Source_Id: sql.NullInt64{Int64: int64(trafficSource.ID), Valid: trafficSource.ID != 0},
	}

	err = db.Where(NewPublisherUrl{
		Publisher_Id:      sql.NullInt64{Int64: int64(publisher.ID), Valid: url.ID != 0},
		Url_Id:            sql.NullInt64{Int64: int64(url.ID), Valid: url.ID != 0},
		Traffic_Source_Id: sql.NullInt64{Int64: int64(trafficSource.ID), Valid: trafficSource.ID != 0},
	}).FirstOrCreate(&newPublisherUrl).Error

	return tagResult, nil
}
