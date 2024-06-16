package db_test

import (
	"testing"
	"time"

	"github.com/clemensjur/sudoku/db"
	"github.com/clemensjur/sudoku/models"
)

func TestWipeDatabase(t *testing.T) {

	db := db.Db("test.db")

	err := db.Migrator().DropTable(&models.Device{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&models.Device{})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDevice(t *testing.T) {
	device := models.Device{
		Name:         "ABC",
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

	db := db.Db("test.db")

	tx := db.Create(&device)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}

func TestReadDevice(t *testing.T) {
	var device models.Device

	db := db.Db("test.db")

	tx := db.First(&device, 1)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}

func TestUpdateDevice(t *testing.T) {
	var device models.Device

	db := db.Db("test.db")

	tx := db.First(&device, 1)
	if tx.Error != nil {
		t.Error(tx.Error)
	}

	device.Name = "test2"

	tx = db.Save(&device)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}

func TestDeleteDevice(t *testing.T) {
	var device models.Device

	db := db.Db("test.db")

	tx := db.First(&device, 1)
	if tx.Error != nil {
		t.Error(tx.Error)
	}

	tx = db.Delete(&device)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}
