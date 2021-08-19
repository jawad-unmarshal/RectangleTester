package RectangleTester

type Rectangle struct {
	length  float64
	breadth float64
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
