package model

type DataType interface {
	Accept(visitor DataTypeVisitor) interface{}
}

type DataTypeVisitor interface {
	VisitText() interface{}
	VisitInteger() interface{}
	VisitReal() interface{}
	VisitBoolean() interface{}
}

type TextDataType struct{}
type IntegerDataType struct{}
type RealDataType struct{}
type BooleanDataType struct{}

func (dt TextDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitText()
}

func (dt IntegerDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitInteger()
}

func (dt RealDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitReal()
}

func (dt BooleanDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitBoolean()
}
