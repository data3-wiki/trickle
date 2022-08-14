package sqlite

import (
	"github.com/dereference-xyz/trickle/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Driver struct {
	dialector gorm.Dialector
}

func NewDriver(dataSourceName string) *Driver {
	return &Driver{dialector: sqlite.Open(dataSourceName)}
}

func (drv *Driver) Dialector() gorm.Dialector {
	return drv.dialector
}

func (drv *Driver) Serialize(dataType model.DataType, value interface{}) interface{} {
	return dataType.Accept(&serializeVisitor{value: value})
}

func (drv *Driver) Deserialize(dataType model.DataType, value interface{}) interface{} {
	return dataType.Accept(&deserializeVisitor{value: value})
}

type serializeVisitor struct {
	value interface{}
}

func (v *serializeVisitor) VisitText() interface{} {
	return v.value
}

func (v *serializeVisitor) VisitInteger() interface{} {
	return v.value
}

func (v *serializeVisitor) VisitReal() interface{} {
	return v.value
}

func (v *serializeVisitor) VisitBoolean() interface{} {
	// TODO
	return v.value
}

type deserializeVisitor struct {
	value interface{}
}

func (v *deserializeVisitor) VisitText() interface{} {
	return v.value
}

func (v *deserializeVisitor) VisitInteger() interface{} {
	return v.value
}

func (v *deserializeVisitor) VisitReal() interface{} {
	return v.value
}

func (v *deserializeVisitor) VisitBoolean() interface{} {
	// TODO
	return v.value
}
