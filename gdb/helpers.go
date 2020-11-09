package gdb

import (
	"errors"

	"gorm.io/gorm"

	"github.com/872409/gatom/log"
)

// import "gorm.io/gorm"

func First(db *gorm.DB, selectFor interface{}) (found interface{}, err error) {
	if err := db.First(selectFor).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorln("repositories RecordNotFound", selectFor)
		return nil, err
	}

	return selectFor, nil
}
