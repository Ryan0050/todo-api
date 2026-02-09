package entity

import "time"

type RecordType string

const (
	NEW        RecordType = "NEW"
	UPDATE     RecordType = "UPDATE"
	DELETE     RecordType = "DELETE"
	ACTIVATE   RecordType = "ACTIVATE"
	DEACTIVATE RecordType = "DEACTIVATE"
)

type Audit struct{
	CreatedAt	time.Time	`gorm:"autoCreatedTime" json:"created_at"`
	CreatedBy	string		`gorm:"type:varchar(100)" json:"created_by"`
	UpdatedAt	time.Time	`gorm:"autoCreatedTime" json:"updated_at"`
	UpdatedBy	string		`gorm:"type:varchar(100)" json:"updated_by"`
	RecType		RecordType	`gorm:"type:varchar(20)" json:"record_type"`
}