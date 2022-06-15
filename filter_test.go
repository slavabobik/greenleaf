package greenleaf

import (
	"reflect"
	"testing"
)

func TestEq(t *testing.T) {
	got := Filter(
		Eq("name", "Slava"),
	)

	want := FilterDocument{
		"name": M{"$eq": "Slava"},
	}

	if !reflect.DeepEqual(want, got) {
		t.FailNow()
	}
}
