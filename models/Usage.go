package models

type Usage struct {
	IdPemakaian 		string  `gorm:"column:id_pemakaian;type:varchar(20);primaryKey" json:"id_pemakaian"`
	CreatedAt           string 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt           string 	`gorm:"column:updated_at" json:"updated_at"`
	IdRuang           string  `gorm:"column:id_ruangan;primaryKey" json:"id_ruangan"`
	AssetCode 		 	string  `gorm:"column:kode_aset;primaryKey" json:"kode_aset"`
	EmployeeID 			string   `gorm:"column:nomor_induk;primaryKey" json:"nomor_induk"`
	Status 				string `gorm:"column:status;" json:"status"`

	Room Room `gorm:"foreignKey:IdRuang" json:"Room"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"Employee"`
	Inventory  Inventory `gorm:"foreignKey:AssetCode" json:"Inventory"`
}

func (Usage) TableName() string {
	return "pemakaian"
}
