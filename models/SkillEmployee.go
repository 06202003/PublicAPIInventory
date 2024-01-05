package models

type SkillEmployee struct {
	SkillEmployeeId string `gorm:"column:id_skill_karyawan;type:varchar(6);primaryKey" json:"id_skill_karyawan"`

	SkillId string `gorm:"column:id_skill;type:varchar(10);" json:"id_skill"`
	Skill   Skill  `gorm:"foreignKey:SkillId;references:SkillId" json:"Skill"`

	EmployeeID string   `gorm:"column:nomor_induk;" json:"nomor_induk"`
	Employee   Employee `gorm:"foreignKey:EmployeeID;references:EmployeeID" json:"Employee"`
}

func (SkillEmployee) TableName() string {
	return "keahlian_karyawan"
}
