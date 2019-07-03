package eziql

// https://www.postgresql.org/docs/11/ddl-constraints.html

const (
	EzrIsNot = "IS NOT"
	EzrEqual = "="
	EzrGreaterThan = ">"
	EzrLessThan = "<"
)

// EzCon holds the details of a constraint on a column in a table
type EzCon struct {
	name string
	column string  // which column is this associated with?
	syntax string
}

func (ec *EzCol) AddCheckConstraint(name, relationship, value string) {
	econ := EzCon{
		name: name,
		column: ec.name,
		syntax: "CONSTRAINT " + name + " CHECK (" + ec.name + " " + relationship + " " + value + ")",
	}
	ec.constraints = append(ec.constraints, econ)
	// "CONSTRAINT positive_price CHECK (price > 0)"
}
