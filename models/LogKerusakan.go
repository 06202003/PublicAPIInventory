package models

type LogKerusakan struct {
	Id          uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IdHistory   string `gorm:"column:id_history" json:"id_history"`
	IdPemakaian string `gorm:"column:id_pemakaian;type:varchar(20);" json:"id_pemakaian"`
	DamageDate  string `gorm:"column:tanggal_kerusakan;type:varchar(20);" json:"tanggal_kerusakan"`
	Usage       Usage  `gorm:"foreignKey:IdPemakaian;references:IdPemakaian" json:"usage"`
}

func (LogKerusakan) TableName() string {
	return "log_kerusakan"
}
