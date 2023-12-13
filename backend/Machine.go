package entity

import (
	"time"

	"gorm.io/gorm"
)

type Machine struct{
	gorm.Model
	M_name string
	Date time.Time
	Pic string `gorm:"type:longtext"`

	
	TypeEquipID *uint
	TypeEquip TypeEquip `gorm:"foreignKey:TypeEquipID"`
	TimeForEquipID *uint
	TimeForEquip TimeForEquip `gorm:"foreignKey:TimeForEquipID"`
}