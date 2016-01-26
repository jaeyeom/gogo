package visitor

func Example() {
	car := Car{
		Wheel("front left"),
		Wheel("front right"),
		Wheel("back left"),
		Wheel("back right"),
		Body{},
		Engine{},
	}
	car.Accept(CarElementPrintVisitor{})
	car.Accept(CarElementDoVisitor{})
	// Output:
	// Visiting front left wheel.
	// Visiting front right wheel.
	// Visiting back left wheel.
	// Visiting back right wheel.
	// Visiting body
	// Visiting engine
	// Visiting car
	// Kicking my front left wheel.
	// Kicking my front right wheel.
	// Kicking my back left wheel.
	// Kicking my back right wheel.
	// Moving my body
	// Starting my engine
	// Starting my car
}
