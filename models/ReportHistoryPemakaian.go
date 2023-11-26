package models



type ReportHistoryPemakaian struct {
	IdHistory               int64       `gorm:"column:id;autoIncrement" json:"id"`
	OldEmployeeID    string    `gorm:"column:nomor_induk_old;type:varchar(45)" json:"nomor_induk_old"`
	NewEmployeeID    string    `gorm:"column:nomor_induk_new;type:varchar(45)" json:"nomor_induk_new"`
	UsageDate        string`gorm:"column:tanggal" json:"tanggal"`
	OldRoom          string    `gorm:"column:ruangan_old;type:varchar(20)" json:"ruangan_old"`
	NewRoom          string    `gorm:"column:ruangan_new;type:varchar(20)" json:"ruangan_new"`
	CreatedAt        string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        string `gorm:"column:updated_at" json:"updated_at"`
	
    IdPemakaian          string `gorm:"column:id_pemakaian;type:varchar(20);primaryKey" json:"id_pemakaian"`
    Usage                Usage  `gorm:"foreignKey:IdPemakaian" json:"Usage"`
}

func (ReportHistoryPemakaian) TableName() string {
	return "history_pemakaian"
}

