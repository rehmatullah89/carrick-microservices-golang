package models

import (
	"database/sql"
	"time"
)

type NewClicksTracking struct {
	ID                     uint          	`gorm:"primarykey"`
	// Publisher_Hash         string         	`gorm:"size:32;not null;index"`
	Publisher_Id      	   uint 			`gorm:"default:null"`
	Click_Identifier_Type  sql.NullString	`gorm:"size:10"`
	Click_Identifier_Value sql.NullString	`gorm:"size:100"`
	Tag                    string        	`gorm:"size:50;not null;index"`
	Publisher_Url          string       	`gorm:"size:2500;comment:Current page url"`
	Click_Url              string         	`gorm:"size:2500"`
	Traffic_Source_Url     sql.NullString	`gorm:"size:2500"`
	Url_Id           	   sql.NullInt64 	`gorm:"default:null"`
	Traffic_Source_Id 	   sql.NullInt64	`gorm:"default:null"`
	Status 	   			   string			`gorm:"default:waiting"`
	Clicks       		   sql.NullInt64 	`gorm:"default:null"`
	User_Agent             string        	`gorm:"size:400"`
	Device_Type            string        	`gorm:"size:10;index"`
	Is_Amp                 bool          	`gorm:"default:false"`
	Ip                     sql.NullString	`gorm:"size:50;default:null"`
	Created_At             time.Time     	`gorm:"default:current_timestamp;index"`
}

func (NewClicksTracking) TableName() string {
	return "new_clicks_tracking"
}
