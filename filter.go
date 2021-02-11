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
	return f.addSingleValue(field, "$eq", value)
}

// Ne adds $ne selector.
func (f *Filter) Ne(field string, value interface{}) *Filter {
	return f.addSingleValue(field, "$ne", value)
}

// In adds $in selector.
func (f *Filter) In(field string, values ...interface{}) *Filter {
	return f.addSlice(field, "$in", values)
}

// Nin adds $nin selector.
func (f *Filter) Nin(field string, values ...interface{}) *Filter {
	return f.addSlice(field, "$nin", values)
}

// Gt adds $gt selector.
func (f *Filter) Gt(field string, value interface{}) *Filter {
	return f.addSingleValue(field, "$gt", value)
}

// Gte adds $gte selector.
func (f *Filter) Gte(field string, value interface{}) *Filter {
	return f.addSingleValue(field, "$gte", value)
}

// Lt adds $lt selector.
func (f *Filter) Lt(field string, value interface{}) *Filter {
	return f.addSingleValue(field, "$lt", value)
}

// Lte adds $lte selector.
func (f *Filter) Lte(field string, value interface{}) *Filter {
	return f.addSingleValue(field, "$lte", value)
}

// Exists adds $exists selector.
func (f *Filter) Exists(field string, value bool) *Filter {
	return f.addSingleValue(field, "$exists", value)
}

func (f *Filter) addSingleValue(field string, operator string, value interface{}) *Filter {
	v, ok := f.selector[field]
	if !ok {
		f.selector[field] = M{operator: value}
		return f
	}

	v[operator] = value
	return f
}

func (f *Filter) addSlice(field, operator string, values []interface{}) *Filter {
	v, ok := f.selector[field]
	if !ok {
		f.selector[field] = M{operator: values}
		return f
	}

	v[operator] = values
	return f
}

// Exec returns document for using in mongodb find operations.
func (f *Filter) Exec() Document {
	return f.selector
}
