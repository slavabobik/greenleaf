package greenleaf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestPagination(t *testing.T) {

	skip := int64(0)
	limit := int64(10)

	pagination := Pagination{
		Limit: &limit,
		Skip:  &skip,
		Sort: Sort{
			"name": DESC,
		},
	}
	got := PaginationOptions(pagination)

	want := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort: Sort{
			"name": DESC,
		},
	}

	assert.Equal(t, want, got)
}

func TestPaginationWithoutSkip(t *testing.T) {

	limit := int64(10)

	pagination := Pagination{
		Limit: &limit,
		Sort: Sort{
			"name": DESC,
		},
	}
	got := PaginationOptions(pagination)

	want := &options.FindOptions{
		Limit: &limit,
		Sort: Sort{
			"name": DESC,
		},
	}

	assert.Equal(t, want, got)
}
