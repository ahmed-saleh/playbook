/*
	This package is the base model for all gorm objects
*/
package models

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ahmed-saleh/playbook/config"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Uuid      string
}

func Setup(m *config.Mysql) {

	var err error
	//TODO: clean up this
	dsn := m.User + ":" + m.Password + "@tcp" + "(" + m.Host + ":" + m.Port + ")/" + m.Name + "?" + "parseTime=true&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&User{})
}
