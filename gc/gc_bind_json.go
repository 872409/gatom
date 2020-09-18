package gc

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	jsoniter "github.com/json-iterator/go"
)

var GCBindingJSON = gcJsonBinding{}

// 可以自定义jsoniter配置或者添加插件
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type gcJsonBinding struct{}

func (gcJsonBinding) Name() string {
	return "json"
}

func (gcJsonBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return decodeJSON(req.Body, obj)
}

func (gcJsonBinding) BindBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if binding.EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	// if binding.EnableDecoderDisallowUnknownFields {
	decoder.DisallowUnknownFields()
	// }
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return ValidateBindObj(obj)
}

func ValidateBindObj(obj interface{}) error {
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(obj)
}

func BindGCJSONBody(gc *GContext, obj interface{}) error {
	body := gc.GetBodyString()
	jsonErr := json.Unmarshal([]byte(body), obj)

	if jsonErr != nil {
		return jsonErr
	}
	return ValidateBindObj(obj)

}
