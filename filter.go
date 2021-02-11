package greenleaf

// Filter represents builder for a find operations.
type Filter struct {
	selector Document
}

// NewFilter returns a *Filter with empty container of find selector.
func NewFilter() *Filter {
	return &Filter{
		selector: make(Document),
	}
}

// Eq adds $eq selector.
func (f *Filter) Eq(field string, value interface{}) *Filter {
	return f.addSelector(field, "$eq", value)
}

// Ne adds $ne selector.
func (f *Filter) Ne(field string, value interface{}) *Filter {
	return f.addSelector(field, "$ne", value)
}

// In adds $in operator with int slice values.
func (f *Filter) In(field string, value interface{}) *Filter {
	return f.addSelector(field, "$in", value)
}

// Nin adds $nin selector.
func (f *Filter) Nin(field string, value interface{}) *Filter {
	return f.addSelector(field, "$nin", value)
}

// Gt adds $gt selector.
func (f *Filter) Gt(field string, value interface{}) *Filter {
	return f.addSelector(field, "$gt", value)
}

// Gte adds $gte selector.
func (f *Filter) Gte(field string, value interface{}) *Filter {
	return f.addSelector(field, "$gte", value)
}

// Lt adds $lt selector.
func (f *Filter) Lt(field string, value interface{}) *Filter {
	return f.addSelector(field, "$lt", value)
}

// Lte adds $lte selector.
func (f *Filter) Lte(field string, value interface{}) *Filter {
	return f.addSelector(field, "$lte", value)
}

// Exists adds $exists selector.
func (f *Filter) Exists(field string, value bool) *Filter {
	return f.addSelector(field, "$exists", value)
}

func (f *Filter) addSelector(field string, operator string, value interface{}) *Filter {
	v, ok := f.selector[field]
	if !ok {
		f.selector[field] = M{operator: value}
		return f
	}

	v[operator] = value
	return f
}

// Exec returns document for using in mongodb find operations.
func (f *Filter) Exec() Document {
	return f.selector
}
