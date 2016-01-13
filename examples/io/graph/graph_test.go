package graph

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestWriteTo(t *testing.T) {
	adjList := [][]int{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}
	w := bytes.NewBuffer(nil)
	if err := WriteTo(w, adjList); err != nil {
		t.Error(err)
	}
	expected := "5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n"
	if expected != w.String() {
		t.Logf("expected: %s\n", expected)
		t.Errorf("found: %s\n", w.String())
	}
}

func ExampleReadFrom() {
	r := strings.NewReader("5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n")
	var adjList [][]int
	if err := ReadFrom(r, &adjList); err != nil {
		fmt.Println(err)
	}
	fmt.Println(adjList)
	// Output:
	// [[3 4] [0 2] [3] [2 4] [0]]
}
