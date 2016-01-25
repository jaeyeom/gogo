package function

func ExampleInsertFunc() {
	m := NewMultiSet()
	ReadFrom(r, InsertFunc(m))
}
