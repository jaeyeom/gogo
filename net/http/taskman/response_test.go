package main

import (
	"encoding/json"
	"fmt"

	"github.com/jaeyeom/gogo/task"
)

func Example_ResponseError_MarshalUnmarshal() {
	err := ResponseError{task.ErrTaskNotExist}
	fmt.Println(err)
	b, _ := json.Marshal(err)
	var err2 ResponseError
	json.Unmarshal(b, &err2)
	fmt.Println(err2)
	// Output:
	// {task does not exist}
	// {task does not exist}
}
