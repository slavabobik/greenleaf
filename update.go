package greenleaf

// Update represents builder for an update queries.
type Update struct {
	operations Document
}

// NewUpdate returns a *Update with empty container of update operations.
func NewUpdate() *Update {
	return &Update{
		operations: make(Document),
	}
}

// SetValue adds $set operator.
// For array values use SetValues method.
func (u *Update) SetValue(field string, value interface{}) *Update {
	return u.addOperator("$set", field, value)
}

// Unset adds $unset operator.
func (u *Update) Unset(field string) *Update {
	return u.addOperator("$unset", field, "")
}

func (u *Update) addOperator(operator, field string, value interface{}) *Update {
	op, ok := u.operations[operator]
	if !ok {
		u.operations[operator] = M{field: value}
		return u
	}
	op[field] = value
	return u
}

// Exec returns document for using in mongodb update operations.
func (u *Update) Exec() Document {
	return u.operations
}
