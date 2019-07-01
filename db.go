// eziql pronounced "easy QL" is a golang ORM that wraps the base database/sql
// package. It is a pure golang implementation that aims to eliminate reflection,
// be very fast and intuitive, and simple.
package eziql

import (
	"strings"
	// "fmt"
	"errors"
	"database/sql"
)

var columns = ["uid", "username", "departname", "created"]
var dtypes = ["serial NOT NULL", "character varying(500) NOT NULL", "character varying(500) NOT NULL", "date"]

type EzDB struct {
	driver string
	dbInfo string
	database *sql.DB
	tables *[]EzTable
}

func New(driver, username, password, dbname, sslmode string) EzDB {
	info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
	edb := EzDB{
		driver: driver,
		dbInfo: info,
	}
	return edb
}

func (edb *EzDB) Connect() error {

}

func (edb *EzDB) AddTable() {

}



func (et *EzTable) Create() error {
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

func (et *EzTable) Insert(data []interface{}) string {
	sqlQuery = "INSERT INTO userinfo(id, name, lastname, other) VALUE($1, $2, $3) returning uid;"
}
