package polymorph

import "fmt"

func ExampleTotalArea() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
	}))
	// Output: 50
}

func ExampleRectangleCircum() {
	r := RectangleCircum{Rectangle{3, 4}}
	fmt.Println(r.Area())
	fmt.Println(r.Circum())
	// Output:
	// 12
	// 14
}

func ExampleWrongRectangle() {
	r := WrongRectangle{Rectangle{3, 4}}
	fmt.Println(r.Area())
	// Output: 24
}

func ExampleTotalArea_moreTypes() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
		RectangleCircum{Rectangle{8, 9}},
		WrongRectangle{Rectangle{1, 2}},
	}))
	// Output: 126
}
