package eziql

import (
	"fmt"
	"strings"
)

const (
	EztBoolean = "boolean"

	// Text types
	EztChar = "char"  // user defined limit char(n)
	EztVarchar = "varchar"  // takes a limit e.g. varchar(n)
	EztText = "text"

	// Numeric types
	EztUint16 = "smallint"
	EztUint32 = "integer"
	EztUint64 = "bigint"
	EztDecimal = "decimal"  // user defined precision
	EztNumeric = "numeric"  // user defined precision (precision, scale)
	EztReal32 = "real"
	EztDouble64 = "double precision"
	EztSerial16 = "smallserial"
	EztSerial32 = "serial"
	EztSerial64 = "bigserial"

	// Date and Time
	EztTimestamp = "timestamp"  // with or without timezone and precision
	EztDate = "date"
	EztTime = "time"  // with or without timezone and precision

	// Qualifiers
	EztTimezone = "with time zone"
)

type EzCol struct {
	name string
	ezt string  // ez type
	precision int
	scale int
	qualifiers []string  // e.g. "with time zone"
	constraints []EzCon
}

func (ec EzCol) CreateColStr() string {
	sqlStr := ec.name + " " + ec.ezt
	if ec.precision >= 0 {
		if ec.scale >= 0 {
			sqlStr += "(" + fmt.Sprintf("%v", ec.precision) + ", " + fmt.Sprintf("%v", ec.scale) + ")"
		} else {
			sqlStr += "(" + fmt.Sprintf("%v", ec.precision) + ")"
		}
	}
	for _, qual := range ec.qualifiers {
		sqlStr += qual + " "
	}
	for _, con := range ec.constraints {
		sqlStr += con.CreateConStr()
	}
	return sqlStr  // "name varchar(150) NOT NULL"
}
