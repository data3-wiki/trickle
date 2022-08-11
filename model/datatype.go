package model

import "fmt"

type DataType string

const (
	Text    DataType = "Text"
	Integer DataType = "Integer"
	Real    DataType = "Real"
)

func (dt DataType) ZeroValue() (interface{}, error) {
	var ret interface{}
	switch dt {
	case Text:
		ret = ""
	case Integer:
		ret = 0
	case Real:
		ret = 0.0
	default:
		return nil, fmt.Errorf("Unsupported DataType: %s", dt)
	}

	return ret, nil
}
