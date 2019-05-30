package main

import (
	"encoding/json"
	"fmt"

	"github.com/jaeyeom/gogo/task"
)

func ExampleResponseError_marshalUnmarshal() {
	err := ResponseError{task.ErrTaskNotExist}
	fmt.Println(err)
	b, _ := json.Marshal(err)
	var err2 ResponseError
	_ = json.Unmarshal(b, &err2)
	fmt.Println(err2)
	// Output:
	// {task does not exist}
	// {task does not exist}
}
