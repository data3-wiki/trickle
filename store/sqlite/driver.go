package sqlite

import (
	"github.com/dereference-xyz/trickle/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite driver that handles sqlite-specific data type conversions.
type Driver struct {
	// Gorm dialector.
	dialector gorm.Dialector
}

// Create new driver with the given dsn string.
func NewDriver(dataSourceName string) *Driver {
	return &Driver{dialector: sqlite.Open(dataSourceName)}
}

// Return gorm dialector.
func (drv *Driver) Dialector() gorm.Dialector {
	return drv.dialector
}

// Serialize the given value based on the data type.
func (drv *Driver) Serialize(dataType model.DataType, value interface{}) interface{} {
	return dataType.Accept(&serializeVisitor{value: value})
}

// Deserialize the given value based on the data type.
func (drv *Driver) Deserialize(dataType model.DataType, value interface{}) interface{} {
	return dataType.Accept(&deserializeVisitor{value: value})
}

// Visitor implementing serialization logic that cases on the data type.
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
	return v.value
}

// Visitor implementing deserialization logic that cases on the data type.
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
	// Note the importance of the decimal point --
	// otherwise the equality might fail because of a type mismatch.
	if v.value == 0.0 {
		return false
	}

	return true
}
