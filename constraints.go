package eziql

type EzCon struct {
	columnName string  // which column is this associated with?
	syntax string
}

func (ec EzCon) CreateConStr() string {
	return "CHECK (column_name IS NOT NULL)"  // "CONSTRAINT table_pkey PRIMARY KEY (uid)"
}
