package decode

import (
	"encoding/json"

	"github.com/dereference-xyz/trickle/model"
)

type convertDecodedVisitor struct {
	value interface{}
	err   error
}

func convertDecodedValue(dataType model.DataType, value interface{}) (interface{}, error) {
	visitor := &convertDecodedVisitor{value: value}
	converted := dataType.Accept(visitor)
	return converted, visitor.err
}

func (v *convertDecodedVisitor) VisitText() interface{} {
	switch v.value.(type) {
	case string:
		return v.value
	default:
		serialized, err := json.Marshal(v.value)
		v.err = err
		return string(serialized)
	}
}

func (v *convertDecodedVisitor) VisitInteger() interface{} {
	return v.value
}

func (v *convertDecodedVisitor) VisitReal() interface{} {
	return v.value
}

func (v *convertDecodedVisitor) VisitBoolean() interface{} {
	return v.value
}
