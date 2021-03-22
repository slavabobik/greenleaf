package greenleaf

import (
	"reflect"
	"testing"
)

func TestFilterBuilder_Exists(t *testing.T) {
	filter := Filter()
	filter.Exists("foo", true)
	filter.Exists("bar", false)
	filter.Exists("foobar", true)

	got := filter.selector
	want := Document{
		"foo":    M{"$exists": true},
		"bar":    M{"$exists": false},
		"foobar": M{"$exists": true},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
