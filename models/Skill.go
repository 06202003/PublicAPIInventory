package models

type Skill struct {
	SkillId        string `gorm:"column:id_skill;type:varchar(6);primaryKey" json:"id_skill"`
	Name      string `gorm:"column:nama" json:"nama"`
	Level      string `gorm:"column:level" json:"level"`
	Notes      string `gorm:"column:notes" json:"notes"`

	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func (Skill) TableName() string {
	return "keahlian"
}
