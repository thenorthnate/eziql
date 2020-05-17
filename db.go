// eziql pronounced "easy QL" is a golang ORM that wraps the base database/sql
// package. It is a pure golang implementation that aims to eliminate reflection,
// be very fast and intuitive, and simple.
package eziql

import (

	// "fmt"
	"database/sql"
	"errors"
)

// QlDB contains the specifics of the database being connected to
type QlDB struct {
	db  *sql.DB
	tab []QlTable
}

func New(conn *sql.DB) (QlDB, error) {
	err := conn.Ping()
	if err != nil {
		return err
	}
	q := QlDB{
		db: conn,
	}
	return q
}

func (edb *QlDB) CreateAll() error {
	for i, _ := range edb.tables {
		err := edb.Create(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (edb *QlDB) Create(table interface{}) error {
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

func (edb *QlDB) AddTable() {

}
