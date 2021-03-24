package gdb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JSONStrings []string

func (c JSONStrings) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSONStrings) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}


// JSONRawValue is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type JSONRawValue []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m JSONRawValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *JSONRawValue) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

type JSONValue JSONRawValue

func (c JSONValue) Value() (driver.Value, error) {
	if len(c) == 0 {
		return nil, nil
	}
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSONValue) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*c = JSONValue(result)
	return err
}

func (c JSONValue) Unmarshal(out interface{}) error {
	return json.Unmarshal(c, out)
}


// MarshalJSON to output non base64 encoded []byte
func (c JSONValue) MarshalJSON() ([]byte, error) {
	return json.RawMessage(c).MarshalJSON()
}

// UnmarshalJSON to deserialize []byte
func (c *JSONValue) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*c = JSONValue(result)
	return err
}

// GormDataType gorm common data type
func (JSONValue) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (JSONValue) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}


type JSONIntegers []int

func (c JSONIntegers) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSONIntegers) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type JSONInt64s []int64

func (c JSONInt64s) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSONInt64s) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type AppClient struct {
	IP      string `json:"ip"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

func (c AppClient) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *AppClient) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
