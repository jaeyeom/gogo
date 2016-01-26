package polymorph

type Shape interface {
	Area() float32
}

type Square struct {
	Size float32
}

func (s Square) Area() float32 {
	return s.Size * s.Size
}

type Rectangle struct {
	Width, Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

type Triangle struct {
	Width, Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Width * t.Height
}

func TotalArea(shapes []Shape) float32 {
	var total float32
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

type RectangleCircum struct{ Rectangle }

func (r RectangleCircum) Circum() float32 {
	return 2 * (r.Width + r.Height)
}

func NewRectangle(width, height float32) *RectangleCircum {
	return &RectangleCircum{Rectangle{width, height}}
}

type WrongRectangle struct{ Rectangle }

func (r WrongRectangle) Area() float32 {
	return r.Rectangle.Area() * 2
}
