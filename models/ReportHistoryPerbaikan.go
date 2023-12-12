package models

type ReportHistoryPerbaikan struct {
	ID                   int64       `gorm:"column:id_perbaikan;primaryKey" json:"id_perbaikan"`
	RepairDate           string `gorm:"column:tanggal_perbaikan" json:"tanggal_perbaikan"`
	Cost                 int64       `gorm:"column:biaya" json:"biaya"`
	Description          string    `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	DamageDate           string `gorm:"column:tanggal_kerusakan" json:"tanggal_kerusakan"`
	RepairCompletionDate string `gorm:"column:tanggal_selesai_perbaikan" json:"tanggal_selesai_perbaikan"`
	CreatedAt            string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            string `gorm:"column:updated_at" json:"updated_at"`
	RepairPlace			 string `gorm:"column:tempat_perbaikan" json:"tempat_perbaikan"`	

	IdHistoryKerusakan   string                  `gorm:"column:id;primaryKey" json:"id"`
    ReportHistoryKerusakan ReportHistoryKerusakan `gorm:"foreignKey:IdHistoryKerusakan" json:"ReportHistoryKerusakan"`
}

func (ReportHistoryPerbaikan) TableName() string {
	return "history_perbaikan"
}
