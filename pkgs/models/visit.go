package models

import (
	"database/sql"
	"time"
)

type Visit struct {
	ID                     uint           `gorm:"primarykey"`
	Publisher_Hash         string         `gorm:"size:32;not null;index"`
	Publisher_Id           sql.NullInt64  `gorm:"default:null"`
	Click_Identifier_Type  sql.NullString `gorm:"size:10;default:null"`
	Click_Identifier_Value sql.NullString `gorm:"size:100;default:null"`
	Tag                    sql.NullString `gorm:"size:50;default:null;index"`
	Traffic_Source_Url     sql.NullString `gorm:"size:2500;default:null"`
	Traffic_Source_Id      sql.NullInt64  `gorm:"default:null"`
	Publisher_Url          string         `gorm:"size:2500;not null;index;comment:Current page url"`
	Publisher_Url_Id       sql.NullInt64  `gorm:"default:null"`
	User_Agent             string         `gorm:"size:400"`
	Device_Type            string         `gorm:"size:10;index"`
	Is_Amp                 bool           `gorm:"default:false"`
	Ip                     sql.NullString `gorm:"size:50;default:null"`
	CreatedAt              time.Time      `gorm:"default:current_timestamp"`

	Publisher     Publisher     `gorm:"constraint:OnDelete:CASCADE"`
	TrafficSource TrafficSource `gorm:"constraint:OnDelete:CASCADE"`
	PublisherUrl  PublisherUrl  `gorm:"constraint:OnDelete:CASCADE"`
}
