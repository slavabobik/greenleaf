package greenleaf

// FilterBuilder represents builder for a find operations.
type FilterBuilder struct {
	selector Document
}

// Filter returs a new instance of a FilterBuilder.
func Filter() *FilterBuilder {
	return &FilterBuilder{
		selector: make(Document),
	}
}

// Eq adds $eq selector.
func (f *FilterBuilder) Eq(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$eq", value)
}

// Ne adds $ne selector.
func (f *FilterBuilder) Ne(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$ne", value)
}

// InInt adds $in operator with int slice values.
func (f *FilterBuilder) InInt(field string, value []int) *FilterBuilder {
	return f.addSelector(field, "$in", value)
}

// InString adds $in operator with string slice values.
func (f *FilterBuilder) InString(field string, value []string) *FilterBuilder {
	return f.addSelector(field, "$in", value)
}

// NinInt adds $nin selector with int slice values.
func (f *FilterBuilder) NinInt(field string, value []int) *FilterBuilder {
	return f.addSelector(field, "$nin", value)
}

// NinString adds $nin selector with int slice values.
func (f *FilterBuilder) NinString(field string, value []string) *FilterBuilder {
	return f.addSelector(field, "$nin", value)
}

// Gt adds $gt selector.
func (f *FilterBuilder) Gt(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$gt", value)
}

// Gte adds $gte selector.
func (f *FilterBuilder) Gte(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$gte", value)
}

// Lt adds $lt selector.
func (f *FilterBuilder) Lt(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$lt", value)
}

// Lte adds $lte selector.
func (f *FilterBuilder) Lte(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$lte", value)
}

// Exists adds $exists selector.
func (f *FilterBuilder) Exists(field string, value bool) *FilterBuilder {
	return f.addSelector(field, "$exists", value)
}

func (f *FilterBuilder) addSelector(field string, operator string, value interface{}) *FilterBuilder {
	v, ok := f.selector[field]
	if !ok {
		f.selector[field] = M{operator: value}
		return f
	}

	v[operator] = value
	return f
}

// Build returns document for using in mongodb find operations.
func (f *FilterBuilder) Build() Document {
	return f.selector
}
