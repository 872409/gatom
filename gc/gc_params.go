package gc

import (
	"github.com/872409/gatom/strs"
)

func (g *gContent) PostInt(name string, def ...int) (val int) {
	_val, _ := g.GetPostForm(name)
	val = strs.StrToInt(_val, def...)
	return
}

func (g *gContent) ParamInt(name string, def ...int) (val int) {
	_val := g.Params.ByName(name)
	val = strs.StrToInt(_val, def...)
	return
}

func (g *gContent) QueryInt(name string, def ...int) (val int) {
	_val := g.Query(name)
	val = strs.StrToInt(_val, def...)
	return
}

func (g *gContent) ParamBoolean(name string, def ...bool) (val bool) {
	_val := g.Params.ByName(name)
	val = strs.StrToBool(_val, def...)
	return
}

func (g *gContent) BindJSONWithError(obj interface{}, code int) (bool, error) {
	if err := g.ShouldBindJSON(obj); err != nil {
		g.JSONErrorWithCode(code)
		return false, err
	}
	return true, nil
}
