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
