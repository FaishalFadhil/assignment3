package models

import (
	"time"

	"gorm.io/gorm"
)

type Wheater struct {
	Id          uint   `gorm:"PRIMARY_KEY"`
	Water       int32  `gorm:"column:water;integer(11)"`
	Wind        int32  `gorm:"column:wind;integer(11)"`
	WaterStatus string `gorm:"column:water_status;varchar(100)"`
	WindStatus  string `gorm:"column:wind_status;varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
