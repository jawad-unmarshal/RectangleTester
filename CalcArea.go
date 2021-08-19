package RectangleTester

type Rectangle struct {
	length  float64
	breadth float64
}

type Square struct {
	side float64
}

func (rect Rectangle) Area() float64 {
	return rect.breadth * rect.length
}

func (rect Rectangle) Perimeter() float64 {
	return 2. * (rect.length + rect.breadth)
}

func (rect Rectangle) SumAreaPerimeter() float64 {
	return rect.Area() + rect.Perimeter()
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 4. * (s.side)
}

func (s Square) SumAreaPerimeter() float64 {
	return s.Area() + s.Perimeter()
}
