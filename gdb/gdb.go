package gdb

import (
	"github.com/872409/gatom/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabaseOption interface {
	GetDatabaseOption(databaseName string) (option DatabaseOption, found bool)
}

type DatabaseOption struct {
	Driver string
	Source string
}

func GetDatabase(name string, databseOptions IDatabaseOption) *gorm.DB {

	option, found := databseOptions.GetDatabaseOption(name)

	if !found {
		panic("GetDatabase option: " + name + " not found")
	}

	database, err := NewDatabase(option)
	if err != nil {
		panic("GetDatabase init: " + name + ",error:" + err.Error())
	}

	return database
}

func NewDatabase(option DatabaseOption) (db *gorm.DB, err error) {
	driver := option.Driver
	if driver == "" {
		driver = "mysql"
	}

	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(option.Source), &gorm.Config{})
		break
	}

	if err != nil {
		log.Errorln("NewDatabase: error ", err, option)
	} else {
		log.Infoln("NewDatabase: ok", option)
	}

	return
}
