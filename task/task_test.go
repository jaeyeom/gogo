package task

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func ExampleTask_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
		nil,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"DONE","deadline":1439739780}
}

func ExampleTask_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":"DONE","Deadline":1439739780}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Buy Milk
	// DONE
	// 2015-08-16 15:43:00 +0000 UTC
}

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{"Name": "John", "Age": 16})
	fmt.Println(string(b))
	// Output:
	// {"Age":16,"Name":"John"}
}

func ExampleTask_String() {
	fmt.Println(Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
		Priority: 0,
		SubTasks: []Task{{"Wash", DONE, nil, 0, nil}, {"Dry", DONE, nil, 0, nil}},
	})
	// Output: [v] Laundry <nil>
}

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}))
	// Output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Put <nil>
	//     [ ] Detergent <nil>
	//   [ ] Dry <nil>
	//   [ ] Fold <nil>
}
