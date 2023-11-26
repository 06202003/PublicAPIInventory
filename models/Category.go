package models

type Category struct {
	ID        string    `gorm:"column:id_kategori;type:varchar(3);primaryKey" json:"id_kategori"`
	Name      string    `gorm:"column:nama;type:varchar(100)" json:"nama"`
	CreatedAt string     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string     `gorm:"column:updated_at" json:"updated_at"`
}

func (Category) TableName() string {
	return "kategori"
}