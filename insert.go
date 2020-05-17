package eziql

type Record interface {
	Pack(v ...interface{})
	Dump() []interface{}
}

func (qdb *QlDB) Insert(r Record) error {
	vals := r.Dump()
	// SQL here to insert records
	_, err := qdb.database.Exec("INSERT INTO userinfo(id, name, lastname, other) VALUE($1, $2, $3) returning uid;", vals...)
	return err
}
