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

// EzDB contains the specifics of the database being connected to
type EzDB struct {
	Driver string
	User string
	Secret string
	Host string
	DBName string
	SSLMode string
	ConnStr string
	database *sql.DB
	tables *[]EzTable
}

func (edb *EzDB) CompileConnStr() {
	if len(edb.Host) > 0 {
		// there is a remote URL specified in for the database
		// "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
		edb.ConnStr = strings.Join([]string{
			edb.Driver,
			"://",
			edb.User,
			":",
			edb.Secret,
			"@",
			edb.Host,
			"/",
			edb.DBName,
		}, "")
		if len(edb.SSLMode) > 0 {
			// ssl mode provided
			edb.ConnStr += "?sslmode" + edb.SSLMode
		}
	} else {
		// no remote URL provided
		edb.ConnStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
	}
}

func (edb *EzDB) Connect() error {
	if len(edb.connStr) == 0 {
		// compile connection string if not already provided
		edb.CompileConnStr()
	}
	db, err := sql.Open(edb.Driver, edb.ConnStr)
	if err != nil {
		return err
	}
	edb.database = db
	return nil
}

func (edb *EzDB) CreateAll() error {
	for i, _ := range edb.tables {
		err := edb.Create(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (edb *EzDB) Create(table interface{}) error {
	if len(edb.tables) == 0 {
		return errors.New("no tables defined for database")
	}
	switch v := table.(type) {
	case int:
		// this should be the index
	case string:
		// this is the table name
	default:
		// nothing specified
	}
	return nil
}

func (edb *EzDB) AddTable() {

}
