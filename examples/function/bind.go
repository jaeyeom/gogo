package function

func InsertFunc(m MultiSet) func(val string) {
	return func(val string) {
		Insert(m, val)
	}
}
