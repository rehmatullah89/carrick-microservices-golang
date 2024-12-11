package models

type PublisherDomain struct {
	ID           uint   `gorm:"primarykey"`
	Publisher_Id uint   `gorm:"not null"`
	Domain       string `gorm:"size:100;not null"`

	Publisher Publisher `gorm:"constraint:OnDelete:CASCADE"`
}
