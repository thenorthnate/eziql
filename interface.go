package eziql

type QlSchema interface {
	Contains() []QlCol
	Pack(data []interface{})
}

func Make(name string, qs QlSchema) QlTable {
	qt := QlTable{
		name: name,
	}

	contents := qs.Contains()
	for _, item := range contents {
		col := buildCol(item)
		qt.columns = append(qt.columns, col)
	}
}
