// Package nocase provides an example code for case insensitive heap.
package nocase

import "strings"

// CaseInsensitive is a slice of string that is sorted case insensitively.
type CaseInsensitive []string

// Len returns the length of c.
func (c CaseInsensitive) Len() int {
	return len(c)
}

// Less returns true if the ith element should sort before the jth
// element.
func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

// Swap swaps the elements with indexes i and j.
func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Push pushes the element x onto the heap.
func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

// Pop removes the minimum element (case insensitively) from the heap
// and returns it.
func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}
