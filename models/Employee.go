package models



type Employee struct {
	EmployeeID string    `gorm:"column:nomor_induk;type:varchar(20);primaryKey" json:"nomor_induk"`
	ImageURL   string    `gorm:"column:gambar;type:varchar(255)" json:"gambar"`
	Name       string    `gorm:"column:nama;type:varchar(100)" json:"nama"`
	Gender     bool      `gorm:"column:gender" json:"gender"`
	Email      string    `gorm:"column:email;type:varchar(100)" json:"email"`
	Phone      string    `gorm:"column:telepon;type:varchar(20)" json:"telepon"`
	Position   string    `gorm:"column:jabatan;type:varchar(20)" json:"jabatan"`
	Division   string    `gorm:"column:divisi;type:varchar(20)" json:"divisi"`
	Address    string    `gorm:"column:alamat;type:varchar(255)" json:"alamat"`
	CreatedAt  string     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  string     `gorm:"column:updated_at" json:"updated_at"`
}

func (Employee) TableName() string {
	return "karyawan"
}
