package eziql

import (
	"strings"
	"errors"
)

// EzTable contains unexported fields that represent the meta data for a SQL table
type EzTable struct {
	name string
	columns []string
	dtypes []string
	uIdx int  // index of the column in the columns field
}

// NewTable returns the struct that contains necessary fields to create and manipulate a database table
func NewTable(tableName, uidCol string, columns, dtypes []string) (EzTable, error) {
	if len(columns) != len(dtypes) {
		return EzTable{}, errors.New("columns and dtypes must be the same length")
	}
	if len(tableName) == 0 {
		return EzTable{}, errors.New("tableName cannot be null")
	}

	uColIdx := -1
	for i := 0; i < len(columns); i++ {
		if column == uidc {
			uColIdx = i
		}
	}
	if uColIdx == -1 && len(uidCol) > 0 {
		return EzTable{}, errors.New("uidCol must match one of the given column names")
	}

	et := EzTable{
		name: tableName,
		columns: columns,
		dtypes: dtypes,
		uIdx: uColIdx,
	}
	return et, nil
}

// AddUIDCol adds a uid column into the table for you
func (et *EzTable) AddUIDCol() {

}

// AddTimeCol adds a time_created column into the table for you
func (et *EzTable) AddTimeCol() {

}
