package table

type Column struct {
	T    Type
	Name string
}

type Schema struct {
	Columns []Column
	Name    string
}
