package models

type PortfolioEmployee struct {
	PortfolioEmployeeId string `gorm:"column:id_portfolio_karyawan;type:varchar(6);primaryKey" json:"id_portfolio_karyawan"`

	PortfolioId string    `gorm:"column:id_portfolio;type:varchar(10);" json:"id_portfolio"`
	Portfolio   Portfolio `gorm:"foreignKey:PortfolioId;references:PortfolioId" json:"Portfolio"`

	EmployeeID string   `gorm:"column:nomor_induk;primaryKey" json:"nomor_induk"`
	Employee   Employee `gorm:"foreignKey:EmployeeID;references:EmployeeID" json:"Employee"`
}

func (PortfolioEmployee) TableName() string {
	return "portfolio_karyawan"
}
