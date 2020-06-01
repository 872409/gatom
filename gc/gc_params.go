package gc

import (
	"errors"
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v8"

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

func (g *gContent) QueryBool(name string, def ...bool) (val bool) {
	_val, exists := g.GetQuery(name)

	if !exists {
		if len(def) > 0 {
			return def[0]
		}
		return false
	}

	val = util.StrToBool(_val, def...)
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

func (g *gContent) BindJSONWithError(obj interface{}) (returnErr error) {
	if bindErr := g.ShouldBindJSON(obj); bindErr != nil {

		for _, e := range bindErr.(validator.ValidationErrors) {
			// fmt.Println("err", k, e.Field, e.Tag, e.Value)
			code, err := strconv.Atoi(e.Name)
			if err != nil {
				code = 1001
			}
			g.JSONErrorWithCodeMsg(code, e.Tag)
			return errors.New(e.Tag)
		}

		return bindErr
	}
	return nil
}
