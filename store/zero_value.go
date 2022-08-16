package store

import "github.com/dereference-xyz/trickle/model"

// Visitor used to case on data type and return the corresponding golang zero value.
type zeroValueVisitor struct{}

// Return the golang zero value corresponding to the given data type.
// This is used by the dynamic-struct library.
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
