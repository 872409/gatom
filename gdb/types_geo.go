package gdb

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

)

type GeoPoint struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (loc GeoPoint) GormDataType() string {
	return "geometry"
}

func (loc GeoPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	// log.Infoln("GormValue", loc)
	return clause.Expr{
		SQL:  "ST_GeomFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%v %v)", loc.Lng, loc.Lat)},
	}
}

func (loc *GeoPoint) Scan(input interface{}) error {
	// log.Infoln("input 1", string(input.([]byte)))

	val := string(input.([]byte))
	if l := len(val); l > 0 {
		values := strings.Split(val[strings.Index(val, "(")+1:len(val)-1], " ")
		if len(values) == 2 {
			if v, err := strconv.ParseFloat(values[0], 64); err == nil {
				loc.Lng = v
			}
			if v, err := strconv.ParseFloat(values[1], 64); err == nil {
				loc.Lat = v
			}
			// log.Infoln("Scan 1", loc)
		}
	}
	return nil
}
