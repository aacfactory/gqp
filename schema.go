package gqp

type Fragment struct {
	Name   string
	OnType string
	Fields []*Field
}

type Aggregation struct {
	Name      string
	FieldName string
}

type Field struct {
	Name        string
	Type        Type
	Query       *Query
	Fragment    *Fragment
	Aggregation *Aggregation
}

type Type struct {
	Name       string
	ArrayTyped bool
}

type QueryCond struct {
}

type QueryRange struct {
	First string
	Size  string
}

type QueryOrder struct {
	FieldName string
	Desc      bool
}

type QueryGroupBy struct {
	FieldNames []string
}

type Query struct {
	Name          string
	Cond          *QueryCond
	Range         *QueryRange
	Orders        []*QueryOrder
	GroupBy       *QueryGroupBy
	Result        *Type
	SelectionSets []*Field
}
