package function

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/jaeyeom/gogo/generate/stringset"
)

func TestInsertFunc(t *testing.T) {
	m := stringset.NewStringSet()
	r := bytes.NewBufferString("a\nb\nc\nc\n")
	ReadFrom(r, InsertFunc(m))
	expected := stringset.StringSet(map[string]int{
		"a": 1,
		"b": 1,
		"c": 2,
	})
	if !reflect.DeepEqual(expected, m) {
		t.Errorf("%d", len(m))
		t.Errorf("%v != %v", expected, m)
	}
}
