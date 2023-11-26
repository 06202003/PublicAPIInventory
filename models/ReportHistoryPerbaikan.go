package models

type ReportHistoryPerbaikan struct {
	ID                   int64       `gorm:"column:id;primaryKey" json:"id"`
	RepairDate           string `gorm:"column:tanggal_perbaikan" json:"tanggal_perbaikan"`
	Cost                 int64       `gorm:"column:biaya" json:"biaya"`
	Description          string    `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	DamageDate           string `gorm:"column:tanggal_kerusakan" json:"tanggal_kerusakan"`
	RepairCompletionDate string `gorm:"column:tanggal_selesai_perbaikan" json:"tanggal_selesai_perbaikan"`
	CreatedAt            string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            string `gorm:"column:updated_at" json:"updated_at"`
	RepairPlace			 string `gorm:"column:tempat_perbaikan" json:"tempat_perbaikan"`	

    IdPemakaian          string `gorm:"column:id_pemakaian;type:varchar(20);" json:"id_pemakaian"`
    Usage                Usage  `gorm:"foreignKey:IdPemakaian;references:IdPemakaian" json:"Usage"`
}

func (ReportHistoryPerbaikan) TableName() string {
	return "history_perbaikan"
}
