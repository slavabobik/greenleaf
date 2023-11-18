package greenleaf

import "go.mongodb.org/mongo-driver/mongo/options"

type sortOrder int

const (
	ASC  sortOrder = 1
	DESC sortOrder = -1
)

// Sort represents the sort options
type Sort map[string]sortOrder

// Pagination options.
type Pagination struct {
	Limit *int64
	Skip  *int64
	Sort  Sort
}

// PaginationOptions returns the MongoDB options for pagination.
func PaginationOptions(page Pagination) *options.FindOptions {
	var options options.FindOptions
	if page.Limit != nil {
		options.Limit = page.Limit
	}

	if page.Skip != nil {
		options.Skip = page.Skip
	}

	if len(page.Sort) > 0 {
		options.Sort = page.Sort
	}

	return &options
}
