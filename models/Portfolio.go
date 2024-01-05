package models

type Portfolio struct {
	PortfolioId string `gorm:"column:id_portfolio;type:varchar(6);primaryKey" json:"id_portfolio"`
	Name        string `gorm:"column:nama" json:"nama"`
	Date        string `gorm:"column:tanggal;type:varchar(20);" json:"tanggal"`
	Deskripsi   string `gorm:"column:deskripsi" json:"deskripsi"`
	Role        string `gorm:"column:role" json:"role"`
	Technology  string `gorm:"column:technology" json:"technology"`

	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func (Portfolio) TableName() string {
	return "portfolio"
}
