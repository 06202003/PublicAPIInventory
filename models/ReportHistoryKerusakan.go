package models



type ReportHistoryKerusakan struct {
	IdHistory            string      `gorm:"column:id;primaryKey" json:"id"`
	IdPemakaian 		string  `gorm:"column:id_pemakaian;type:varchar(20);primaryKey" json:"id_pemakaian"`
    DamageDate			 string 	`gorm:"column:tanggal_kerusakan;type:varchar(20);" json:"tanggal_kerusakan"`
	Usage Usage 		`gorm:"foreignKey:IdPemakaian;references:IdPemakaian" json:"Usage"`
}

func (ReportHistoryKerusakan) TableName() string {
	return "history_kerusakan"
}

