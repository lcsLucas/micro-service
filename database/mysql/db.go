package database

import (
	"fmt"

	"github.com/lcslucas/micro-service/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBURL string

func Connect(c config.ConfigDB) (*gorm.DB, error) {
	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DBName)

	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
