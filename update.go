package greenleaf

// UpdateBuilder represents builder for an update queries.
type UpdateBuilder struct {
	operations Document
}

// Update returs a new instance of a UpdateBuilder.
func Update() *UpdateBuilder {
	return &UpdateBuilder{
		operations: make(Document),
	}
}

// Set adds $set operator.
func (u *UpdateBuilder) Set(field string, value interface{}) *UpdateBuilder {
	return u.addOperator("$set", field, value)
}

// Unset adds $unset operator.
func (u *UpdateBuilder) Unset(field string) *UpdateBuilder {
	return u.addOperator("$unset", field, "")
}

func (u *UpdateBuilder) addOperator(operator, field string, value interface{}) *UpdateBuilder {
	op, ok := u.operations[operator]
	if !ok {
		u.operations[operator] = M{field: value}
		return u
	}
	op[field] = value
	return u
}

// Build returns document for using in mongodb update operations.
func (u *UpdateBuilder) Build() Document {
	return u.operations
}
