package model

// Data type of a particular value.
type DataType interface {
	// Used by implementations to invoke the appropriate visit method based on the data type.
	Accept(visitor DataTypeVisitor) interface{}
}

// Visitor interface to be implemented whenever we have logic that has to case on the data type.
// The visitor pattern allows us to use the type system to enforce case coverage of all data types.
// i.e. Poor man's implementation of sum types.
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
