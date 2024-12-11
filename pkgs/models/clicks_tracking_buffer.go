package models

import (
	"database/sql"
	"time"
)

type ClicksTrackingBuffer struct {
	ID                        uint           `gorm:"primarykey"`
	Publisher_Hash            string         `gorm:"size:32;not null;index"`
	Click_Identifier_Type     sql.NullString `gorm:"size:10"`
	Click_Identifier_Value    sql.NullString `gorm:"size:100"`
	Tag                       string         `gorm:"size:50;not null;index"`
	Traffic_Source_Url        sql.NullString `gorm:"size:2500"`
	Publisher_Url             string         `gorm:"size:2500;comment:Current page url"`
	Click_Url                 string         `gorm:"size:2500"`
	User_Agent                string         `gorm:"size:400"`
	Device_Type               string         `gorm:"size:10;index"`
	Is_Amp                    bool           `gorm:"default:false"`
	Ip                        sql.NullString `gorm:"size:50;default:null"`
	Created_At                time.Time      `gorm:"default:current_timestamp;index"`
	Affiliate_Network_Type_Id sql.NullString `gorm:"default:null"`
}
