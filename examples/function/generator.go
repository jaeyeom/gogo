package function

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

type VertexID int
type EdgeID int

func NewVertexIDGenerator() func() VertexID {
	var next int
	return func() VertexID {
		next++
		return VertexID(next)
	}
}
