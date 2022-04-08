package greenleaf

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestEq(t *testing.T) {
	got := Filter(
		Eq("name", "Slava"),
	)

	want := Document{
		"name": M{"$eq": "Slava"},
	}

	if !reflect.DeepEqual(want, got) {
		_, file, line, _ := runtime.Caller(0)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, want, got)
		t.FailNow()
	}
}
