package greenleaf

type operator string

const (
	eqOperator     = operator("$eq")
	neOperator     = operator("$ne")
	gtOperator     = operator("$gt")
	gteOperator    = operator("$gte")
	ltOperator     = operator("$lt")
	lteOperator    = operator("$lte")
	inOperator     = operator("$in")
	existsOperator = operator("$exists")
	ninOperator    = operator("$nin")
)

// FilterDocument represents filter document.
type FilterDocument map[string]M

// M represents selector.
type M map[operator]any

// Eq creates a $eq query selector.
func Eq[T any](field string, value T) FilterDocument {
	return createSelector(field, eqOperator, value)
}

// Ne creates a $ne query selector.
func Ne[T any](field string, value T) FilterDocument {
	return createSelector(field, neOperator, value)
}

// Gt creates a $gt query selector.
func Gt[T any](field string, value T) FilterDocument {
	return createSelector(field, gtOperator, value)
}

// Gte creates a $gte query selector.
func Gte[T any](field string, value T) FilterDocument {
	return createSelector(field, gteOperator, value)
}

// Lt creates a $lt query selector.
func Lt[T any](field string, value T) FilterDocument {
	return createSelector(field, ltOperator, value)
}

// Lte creates a $lte query selector.
func Lte[T any](field string, value T) FilterDocument {
	return createSelector(field, lteOperator, value)
}

// Exists creates a $exists query selector.
func Exists(field string, value bool) FilterDocument {
	return createSelector(field, existsOperator, value)
}

// In creates a $in query selector.
func In[T any](field string, value []T) FilterDocument {
	return createSelector(field, inOperator, value)
}

// Nin creates a $nin query selector.
func Nin[T any](field string, value []T) FilterDocument {
	return createSelector(field, ninOperator, value)
}

func createSelector(field string, operator operator, value any) FilterDocument {
	return FilterDocument{
		field: M{operator: value},
	}
}

func Filter(docs ...FilterDocument) FilterDocument {
	query := FilterDocument{}
	for _, doc := range docs {
		for k, v := range doc {
			query[k] = v
		}
	}

	return query
}
