package device

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Name         string
	Type         string
	Mac          string
	Ip           string
	Ipv6         string
	Socket       string
	Location     string
	Status       string
	Department   string
	Os           string
	BoughtAt     time.Time
	Warranty     time.Time
	Manufacturer string
	Product      string
	Serial       string
	LastSeen     time.Time
}
