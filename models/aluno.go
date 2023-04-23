package models

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `gorm:"size:255; not null;" json:"nome"`
	CPF  string `gorm:"size:11; not null;" json:"cpf"`
	RG   string `gorm:"size:9; not null;" json:"rg"`
}
