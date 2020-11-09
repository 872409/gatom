package gdb

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONStrings []string

func (c JSONStrings) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSONStrings) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
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