package db_test

import (
	"testing"
	"time"

	"github.com/clemensjur/sudoku/db"
	"github.com/clemensjur/sudoku/device"
)

func TestCreateDevice(t *testing.T) {
	device := device.Device{
		Name:         "test",
		Type:         "test",
		Mac:          "test",
		Ip:           "test",
		Ipv6:         "test",
		Socket:       "test",
		Location:     "test",
		Status:       "test",
		Department:   "test",
		Os:           "test",
		BoughtAt:     time.Now(),
		Warranty:     time.Now(),
		Manufacturer: "test",
		Product:      "test",
		Serial:       "test",
		LastSeen:     time.Now(),
	}

	tx := db.Db().Create(&device)

	if tx.Error != nil {
		t.Error(tx.Error)
	}
}

func TestReadDevice(t *testing.T) {
	var device device.Device

	tx := db.Db().First(&device, 1)

	if tx.Error != nil {
		t.Error(tx.Error)
	}
}

func TestUpdateDevice(t *testing.T) {
	var device device.Device
	db.Db().First(&device, 1)
	device.Name = "test2"

	tx := db.Db().Save(&device)

	if tx.Error != nil {
		t.Error(tx.Error)
	}
}

func TestDeleteDevice(t *testing.T) {
	var device device.Device
	db.Db().First(&device, 1)

	tx := db.Db().Delete(&device)

	if tx.Error != nil {
		t.Error(tx.Error)
	}
}
