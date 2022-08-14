package store

import "github.com/dereference-xyz/trickle/model"

type zeroValueVisitor struct{}

func zeroValue(dataType model.DataType) interface{} {
	return dataType.Accept(&zeroValueVisitor{})
}

func (v *zeroValueVisitor) VisitText() interface{} {
	return ""
}

func (v *zeroValueVisitor) VisitInteger() interface{} {
	return 0
}

func (v *zeroValueVisitor) VisitReal() interface{} {
	return 0.0
}

func (v *zeroValueVisitor) VisitBoolean() interface{} {
	return false
}
