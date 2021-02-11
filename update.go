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

// SetValue adds $set operator with single value e.g { "quantity": 500 } .
// For array values use SetValues method.
func (u *Update) SetValue(field string, value interface{}) *Update {
	return u.addSingleValue("$set", field, value)
}

// SetValues adds $set operator with array value e.g  { "tags": [ "coats", "outerwear", "clothing" ] }.
// For single value use SetValue method.
func (u *Update) SetValues(field string, values ...interface{}) *Update {
	return u.addSlice("$set", field, values)
}

// Unset adds $unset operator.
func (u *Update) Unset(field string) *Update {
	return u.addSingleValue("$unset", field, "")
}

func (u *Update) addSingleValue(operator, field string, value interface{}) *Update {
	op, ok := u.operations[operator]
	if !ok {
		u.operations[operator] = M{field: value}
		return u
	}
	op[field] = value
	return u
}

func (u *Update) addSlice(operator, field string, values ...interface{}) *Update {
	op, ok := u.operations[operator]
	if !ok {
		u.operations[operator] = M{field: values}
		return u
	}
	op[field] = values
	return u
}

// Exec returns document for using in mongodb update operations.
func (u *Update) Exec() Document {
	return u.operations
}
