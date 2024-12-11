package models

import (
	"database/sql"
	"time"
)

type PublisherUrl struct {
	ID                uint          `gorm:"primarykey"`
	Publisher_Hash    string        `gorm:"size:32;not null;index"`
	Publisher_Id      sql.NullInt64 `gorm:"default:null"`
	Url_Id            sql.NullInt64 `gorm:"default:null"`
	TagStr            string        `gorm:"column:tag;size:50;not null"`
	Tag_Id            sql.NullInt64 `gorm:"default:null"`
	Traffic_Source_Id sql.NullInt64 `gorm:"default:null"`
	Created_At        time.Time     `gorm:"default:current_timestamp"`

	Publisher     Publisher     `gorm:"constraint:OnDelete:CASCADE"`
	Url           Url           `gorm:"constraint:OnDelete:CASCADE"`
	Tag           Tag           `gorm:"constraint:OnDelete:CASCADE"`
	TrafficSource TrafficSource `gorm:"constraint:OnDelete:CASCADE"`
	Visits        []Visit       `gorm:"constraint:OnDelete:CASCADE"`
}

func (PublisherUrl) TableName() string {
	return "publisher_url"
}
