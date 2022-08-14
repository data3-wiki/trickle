package model

type DataType interface {
	Accept(visitor DataTypeVisitor) interface{}
}

type DataTypeVisitor interface {
	VisitText() interface{}
	VisitInteger() interface{}
	VisitReal() interface{}
}

type TextDataType struct{}
type IntegerDataType struct{}
type RealDataType struct{}

func (dt TextDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitText()
}

func (dt IntegerDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitInteger()
}

func (dt RealDataType) Accept(visitor DataTypeVisitor) interface{} {
	return visitor.VisitReal()
}
