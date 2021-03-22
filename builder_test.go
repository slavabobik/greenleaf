package greenleaf

import (
	"reflect"
	"testing"
)

func TestFilterBuilder_Exists(t *testing.T) {
	filter := Filter()
	filter.Exists("foo", true)
	filter.Exists("bar", false)

	got := filter.selector
	want := Document{
		"foo": M{"$exists": true},
		"bar": M{"$exists": false},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
