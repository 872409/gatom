package gdb

import (
	"github.com/872409/gatom/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseOption struct {
	Driver string
	Source string
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
		log.Errorln("initDBSvr: error ", err)
	} else {
		log.Infoln("initDBSvr: ok")
	}

	return
}
