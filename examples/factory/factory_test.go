package factory

func ExampleCreateFactory() {
	f1 := CreateFactory("win")
	Run(f1)
	f2 := CreateFactory("mac")
	Run(f2)
	// Output:
	// win button paint
	// win button click
	// win label paint
	// mac button paint
	// mac button click
	// mac label paint
}
