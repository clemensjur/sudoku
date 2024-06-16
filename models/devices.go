package models

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Name         string `gorm:"serializer:json"`
	Type         string `gorm:"serializer:json"`
	Mac          string `gorm:"serializer:json"`
	Ip           string `gorm:"serializer:json"`
	Ipv6         string `gorm:"serializer:json"`
	Socket       string `gorm:"serializer:json"`
	Location     string `gorm:"serializer:json"`
	Status       string `gorm:"serializer:json"`
	Department   string `gorm:"serializer:json"`
	Os           string `gorm:"serializer:json"`
	BoughtAt     time.Time
	Warranty     time.Time
	Manufacturer string `gorm:"serializer:json"`
	Product      string `gorm:"serializer:json"`
	Serial       string `gorm:"serializer:json"`
	LastSeen     time.Time
}
