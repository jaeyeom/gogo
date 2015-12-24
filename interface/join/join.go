// Package join provide an example of interface and type switch.
package join

import (
	"fmt"
	"strconv"
	"strings"
)

// Join concatenates the elements of a to create a single string. The
// separator sep is placed between each element.
func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
		// The switch-case block above is equivalent to the
		// following if-else if.
		//
		// if x, ok := a[i].(string); ok {
		// 	t[i] = x
		// } else if x, ok := a[i].(int); ok {
		// 	t[i] = strconv.Itoa(x)
		// } else if x, ok := a[i].(fmt.Stringer); ok {
		// 	t[i] = x.String()
		// }
	}
	return strings.Join(t, sep)
}
