/*
	This package is the base model for all gorm objects
*/
package models

import (
	"fmt"
	"log"
	"time"

	"github.com/ahmed-saleh/playbook/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Uuid      string
}

func Setup() {

	var err error
	//TODO: clean up this
	dsn := config.MysqlSettings.User + ":" + config.MysqlSettings.Password + "@tcp" + "(" + config.MysqlSettings.Host + ":" + config.MysqlSettings.Port + ")/" + config.MysqlSettings.Name + "?" + "parseTime=true&loc=Local"
	fmt.Println(dsn)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&User{})
}
