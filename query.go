package eziql

type EzWhere struct {
	Column   string
	Relation string
	Value    string
}

type QlRecord struct {
	records [][]*QlCol
}

func (ql *QlDB) Select() ([]QlRecord, error) {

}

/*
rows, err := db.Select(
	Cols("id", "name"),
	Where(
		GreaterThan("id", 4),
		Contains("name", "Nate"),
	),
)
*/

func (ql *QlTable) Data(f ...func(v ...interface{})) {
	if len(f) == 0 {
		// query all values in the table!
		qr := ql.Select()
		return qr
	}
}
