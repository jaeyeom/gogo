package join

import (
	"fmt"

	"github.com/jaeyeom/gogo/task"
)

func ExampleJoin() {
	t := task.Task{
		Title:    "Laundry",
		Status:   task.DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	// Output: 1,two,3,[v] Laundry <nil>
}
