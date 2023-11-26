package models

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Status   string `gorm:"type:varchar(100)" json:"status"`
	Nama     string `gorm:"type:varchar(100)" json:"nama"`
}