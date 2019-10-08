package gc

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/872409/gatom/util"
)

func (g *gContent) PostInt(name string, def ...int) (val int) {
	_val, _ := g.GetPostForm(name)
	val = util.StrToInt(_val, def...)
	return
}

func (g *gContent) PostInt64(name string, def ...int) (val int64, err error) {
	_val, exists := g.GetPostForm(name)
	if !exists {
		return 0, errors.New("not exists")
	}

	val, err = strconv.ParseInt(_val, 10, 64)
	if err != nil {
		fmt.Println("QueryInt64", _val, err)
		return 0, errors.New("parse error")
	}

	return val, nil
}

func (g *gContent) ParamInt(name string, def ...int) (val int) {
	_val := g.Params.ByName(name)
	val = util.StrToInt(_val, def...)
	return
}

func (g *gContent) QueryInt(name string, def ...int) (val int) {
	_val := g.Query(name)
	val = util.StrToInt(_val, def...)
	return
}
func (g *gContent) QueryInt64(name string) (val int64, err error) {
	_val, exists := g.GetQuery(name)

	if exists {
		val, err := strconv.ParseInt(_val, 10, 64)
		if err != nil {
			fmt.Println("QueryInt64", val, err)
			return 0, err
		}
		return val, nil
	}

	return 0, errors.New("not exits")
}

func (g *gContent) ParamBoolean(name string, def ...bool) (val bool) {
	_val := g.Params.ByName(name)
	val = util.StrToBool(_val, def...)
	return
}

func (g *gContent) BindJSONWithError(obj interface{}, errDef error) error {
	if err := g.ShouldBindJSON(obj); err != nil {
		g.JSONCodeError(errDef)
		return err
	}
	return nil
}
