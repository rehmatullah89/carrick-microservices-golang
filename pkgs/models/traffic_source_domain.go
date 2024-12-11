package models

import "time"

type TrafficSourceDomain struct {
	ID                uint      `gorm:"primarykey"`
	Traffic_Source_Id uint      `gorm:"not null"`
	Domain            string    `gorm:"size:100;not null,unique"`
	Created_At        time.Time `gorm:"default:current_timestamp"`
	Updated_At        time.Time `gorm:"default:current_timestamp"`

	TrafficSource TrafficSource `gorm:"constraint:OnDelete:CASCADE"`
}
