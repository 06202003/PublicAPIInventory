package models

type Room struct {
	Id        string `gorm:"column:id_ruangan;type:varchar(6);primaryKey" json:"id_ruangan"`
	Name      string `gorm:"column:nama" json:"nama"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`

	IdLokasi string   `gorm:"column:id_lokasi;type:varchar(10);" json:"id_lokasi"`
	Location Location `gorm:"foreignKey:IdLokasi;references:IdLokasi" json:"Location"`
}

func (Room) TableName() string {
	return "ruangan"
}
