package entity

type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Job         string `gorm:"not null" json:"job"`
	Description string `json:"description"`
	Status      string `gorm:"default:on progress" json:"status"`
	Audit       Audit  `gorm:"embedded" json:"audit"`
}
