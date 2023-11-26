package models

type Location struct {
	IdLokasi        string  `gorm:"column:id_lokasi;type:varchar(5);primaryKey" json:"id_lokasi"`
	Name           	string 	`gorm:"column:nama;type:varchar(100)" json:"nama"`
	Alamat          string 	`gorm:"column:alamat;type:varchar(100)" json:"alamat"`
	CreatedAt       string 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt       string 	`gorm:"column:updated_at" json:"updated_at"`
}

func (Location) TableName() string {
	return "lokasi"
}
