package maps

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCount_deepEqual(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func TestCount_if(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if len(codeCount) != 3 {
		t.Error("codeCount:", codeCount)
		t.Fatal("count should be 3 but:", len(codeCount))
	}
	if codeCount['가'] != 1 || codeCount['나'] != 2 || codeCount['다'] != 1 {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func ExampleCount_sort() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
	// Output:
	// 가 1
	// 나 2
	// 다 1
}
