package swagger

import "github.com/dereference-xyz/trickle/model"

type swaggerTypeVisitor struct{}

func swaggerType(dataType model.DataType) string {
	return dataType.Accept(&swaggerTypeVisitor{}).(string)
}

func (v *swaggerTypeVisitor) VisitText() interface{} {
	return "string"
}

func (v *swaggerTypeVisitor) VisitInteger() interface{} {
	return "integer"
}

func (v *swaggerTypeVisitor) VisitReal() interface{} {
	return "number"
}

func (v *swaggerTypeVisitor) VisitBoolean() interface{} {
	return "boolean"
}
