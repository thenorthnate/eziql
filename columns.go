package eziql

// https://www.postgresql.org/docs/11/datatype.html

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

var (
	EzType = map[string]string{
		"uint16": "smallint",
		"uint32": "integer",
		"int": "integer",
		"uint64": "bigint",
		"decimal": "decimal",
		"real": "real",
		"float32": "real",
		"float64": "double precision",
		"serial16": "smallserial",
		"serial32": "serial",
		"serial64": "bigserial",
		"timestamp": "timestamp",
		"date": "date",
		"time": "time",
	}
)

type QlCol struct {
	name string
	ezt string  // ez type
	precision int
	scale int
	qualifiers []string  // e.g. "with time zone"
	constraints []EzCon
}

func (qc QlCol) CreateColStr() string {
	sqlStr := qc.name + " " + qc.ezt
	if qc.precision >= 0 {
		if qc.scale >= 0 {
			sqlStr += ParenWrap(fmt.Sprintf("%v", qc.precision) + ", " + fmt.Sprintf("%v", qc.scale))
		} else {
			sqlStr += ParenWrap(fmt.Sprintf("%v", qc.precision))
		}
	}
	for _, qual := range qc.qualifiers {
		sqlStr += qual + " "
	}
	for _, con := range qc.constraints {
		sqlStr += con.syntax + " "
	}
	return sqlStr  // "name varchar(150) NOT NULL"
}
