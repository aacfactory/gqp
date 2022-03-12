package gqp

type Fragment struct {
	Name   string
	OnType string
	Fields []*Field
}

func (fragment *Fragment) params() (v []string, has bool) {
	if fragment.Fields != nil && len(fragment.Fields) > 0 {
		for _, field := range fragment.Fields {
			fieldParams, fieldHasParams := field.params()
			if fieldHasParams {
				v = append(v, fieldParams...)
			}
		}
	}
	has = len(v) > 0
	return
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

func (field *Field) params() (v []string, has bool) {
	if field.Query != nil {
		queryParams, queryHasParams := field.Query.Params()
		if queryHasParams {
			v = append(v, queryParams...)
		}
	}
	if field.Fragment != nil {
		fragmentParams, fragmentHasParams := field.Fragment.params()
		if fragmentHasParams {
			v = append(v, fragmentParams...)
		}
	}
	has = len(v) > 0
	return
}

type Type struct {
	Name       string
	ArrayTyped bool
}

type QueryCond struct {
	OR        bool
	FieldName string
	Operation string
	Values    []interface{}
	Params    []string
	Children  []*QueryCond
}

func (cond *QueryCond) params() (v []string, has bool) {
	if cond.Params != nil && len(cond.Params) > 0 {
		v = append(v, cond.Params...)
	}
	if cond.Children != nil && len(cond.Children) > 0 {
		for _, child := range cond.Children {
			childParams, childHasParams := child.params()
			if childHasParams {
				v = append(v, childParams...)
			}
		}
	}
	has = len(v) > 0
	return
}

type QueryRange struct {
	First      int
	FirstParam string
	Size       int
	SizeParam  string
}

type QueryOrder struct {
	FieldName string
	Desc      bool
}

type QueryGroupBy struct {
	FieldNames []string
}

type Query struct {
	Name       string
	Cond       *QueryCond
	Range      *QueryRange
	Orders     []*QueryOrder
	GroupBy    *QueryGroupBy
	Result     *Type
	Selections []*Field
}

func (q *Query) Params() (v []string, has bool) {
	// cond
	if q.Cond != nil {
		condParams, condHasParams := q.Cond.params()
		if condHasParams {
			v = append(v, condParams...)
		}
	}
	// Selections
	if q.Selections != nil && len(q.Selections) > 0 {
		for _, selection := range q.Selections {
			selectionParams, selectionHasParams := selection.params()
			if selectionHasParams {
				v = append(v, selectionParams...)
			}
		}
	}
	// Range
	if q.Range != nil {
		if q.Range.FirstParam != "" {
			v = append(v, q.Range.FirstParam)
		}
		if q.Range.SizeParam != "" {
			v = append(v, q.Range.SizeParam)
		}
	}
	has = len(v) > 0
	return
}
