package eziql

import (
	"strings"
	"errors"
)

// EzTable contains unexported fields that represent the meta data for a SQL table
type EzTable struct {
	name string  // table name
	columns []EzCol
	constraints []string
}

// NewTable returns the struct that contains necessary fields to create and manipulate a database table
func NewTable(tableName, columns []EzCol) (EzTable, error) {
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

func (et *EzTable) CreateTableStr() error {
	createArray := []string{}
	for i := 0; i < len(et.columns); i++ {
		createColumnString := strings.Join([]string{et.columns[i], et.dtypes[i]}, " ")
		createArray = append(createArray, createColumnString)
	}
	if et.uIdx != -1 {
		// we want to add in the unique column
		ccS := "CONSTRAINT " + et.name + "_pkey PRIMARY KEY (" + et.columns[et.uIdx] + ")"
		createArray = append(createArray, ccS)
	}
	cStr := strings.Join(createArray, ", ")

	sqlQuery = "CREATE TABLE " + et.name "(" cStr + ")"
}

func (et *EzTable) InsertStr(data map[string]interface{}) string {
	sqlQuery = "INSERT INTO userinfo(id, name, lastname, other) VALUE($1, $2, $3) returning uid;"
}

func (et *EzTable) SelectStr(data map[string]interface{}) string {
	sqlQuery = "SELECT * FROM userinfo;"
}
