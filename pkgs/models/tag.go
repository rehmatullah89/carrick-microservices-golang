package models

import (
	"carrick-js-api/pkgs/db"
	"database/sql"
)

type Tag struct {
	ID           uint           `gorm:"primarykey"`
	Tag          string         `gorm:"size:50;not null;unique"`
	Used         bool           `gorm:"default:false;not null"`
	Traffic_Type sql.NullString `gorm:"size:20;default:null"`
	Publisher_Id uint           `gorm:"not null"`
	db.Timestamps

	Publisher     Publisher      `gorm:"constraint:OnDelete:CASCADE"`
	PublisherUrls []PublisherUrl `gorm:"constraint:OnDelete:CASCADE"`
}

type TagResult struct {
	Id  uint
	Tag string
}

func getFreeTag(publisher Publisher, traffic_type string) Tag {
	db := db.GetDBInstance().GetDB()

	var tag Tag

	db.Raw(`UPDATE tags
		SET    used = true, traffic_type=@traffic_type, updated_at = current_timestamp
		WHERE  id = (
			SELECT id
			FROM   tags
			WHERE  used = false
				AND publisher_id = @publisher_id
			ORDER BY updated_at
			LIMIT  1
		)
		RETURNING *`, map[string]interface{}{
		"traffic_type": traffic_type,
		"publisher_id": publisher.ID,
	}).Scan(&tag)

	return tag
}

func GetTagForTrafficType(publisher Publisher, trafficType string) TagResult {
	tag := getFreeTag(publisher, trafficType)

	if tag.ID == 0 {
		return TagResult{}
	}

	return TagResult{Id: tag.ID, Tag: tag.Tag}
}
