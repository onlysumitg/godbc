package godbc

import (
	"reflect"
	"strings"
)

var TimeFormat string = "15:04:05"
var DateFormat string = "2006-01-02"
var TimestampFormat string = "2006-01-02 15:04:05.000000"

type Col struct {
	Index int `json:"-"`
	Name  string
	Type  reflect.Type `json:"-"`
	Value any
}

type RSRows []Col

func (c *Col) AssignValueByType() {

	switch c.Type.Kind() {
	case reflect.String:
		if c.Value == nil {
			c.Value = nil
		} else {
			c.Value = strings.TrimSpace(asString(c.Value))
		}
	}
}

func (r RSRows) ToMap() map[string]any {
	m := make(map[string]any)

	for _, c := range r {
		c.AssignValueByType()
		m[c.Name] = c.Value
	}

	return m
}
