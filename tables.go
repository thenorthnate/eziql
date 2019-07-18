package eziql

import (
	"strings"
	"errors"
)

// QlTable contains unexported fields that represent the meta data for a SQL table
type QlTable struct {
	name string  // table name
	columns []EzCol
	constraints []string
}

// NewTable returns the struct that contains necessary fields to create and manipulate a database table
func NewTable(tableName, columns []QlCol) (QlTable, error) {
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

	qt := QlTable{
		name: tableName,
		columns: columns,
		dtypes: dtypes,
		uIdx: uColIdx,
	}
	return qt, nil
}

// AddUIDCol adds a uid column into the table for you
func (et *QlTable) AddUIDCol() {

}

// AddTimeCol adds a time_created column into the table for you
func (et *QlTable) AddTimeCol() {

}

func (et *QlTable) CreateTableStr() error {
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

func (et *QlTable) InsertStr(data map[string]interface{}) string {
	sqlQuery = "INSERT INTO userinfo(id, name, lastname, other) VALUE($1, $2, $3) returning uid;"
}

func (et *QlTable) SelectStr(data map[string]interface{}) string {
	sqlQuery = "SELECT * FROM userinfo;"
}
