package models

import (
	"gorm.io/gorm"
)

type Creator struct {
	gorm.Model

	Name string

	Languages []Language `gorm:"many2many:language_creators" json:"-"`
}
