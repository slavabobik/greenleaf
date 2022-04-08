package greenleaf

const (
	eqOp     = "$eq"
	neOp     = "$ne"
	gtOp     = "$gt"
	gteOp    = "$gte"
	ltOp     = "$lt"
	lteOp    = "$lte"
	inOp     = "$in"
	existsOp = "$exists"
	ninOp    = "$nin"
)

func Eq[T comparable](field string, value T) Document {
	return createSelector(field, eqOp, value)
}

func EqSlice[T any](field string, value []T) Document {
	return createSelector(field, eqOp, value)
}

func Ne[T comparable](field string, value T) Document {
	return createSelector(field, neOp, value)
}

func Gt[T comparable](field string, value T) Document {
	return createSelector(field, gtOp, value)
}

func Gte[T comparable](field string, value T) Document {
	return createSelector(field, gteOp, value)
}

func Lt[T comparable](field string, value T) Document {
	return createSelector(field, ltOp, value)
}

func Lte[T comparable](field string, value T) Document {
	return createSelector(field, lteOp, value)
}

func Exists(field string, value bool) Document {
	return createSelector(field, existsOp, value)
}

func In[T any](field string, value []T) Document {
	return createSelector(field, inOp, value)
}

func Nin[T any](field string, value []T) Document {
	return createSelector(field, ninOp, value)
}

func createSelector(field, operator string, value any) Document {
	return Document{
		field: M{operator: value},
	}
}

func Filter(docs ...Document) Document {
	query := Document{}
	for _, doc := range docs {
		for k, v := range doc {
			query[k] = v
		}
	}

	return query
}
