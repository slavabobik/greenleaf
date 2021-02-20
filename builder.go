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

// Exists adds $exists selector.
func (f *FilterBuilder) Exists(field string, value bool) *FilterBuilder {
	return f.addSelector(field, "$exists", value)
}
