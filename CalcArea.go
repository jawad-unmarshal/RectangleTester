package RectangleTester

func RectArea(length, breadth float64) float64 {
	return length * breadth
}
func RectPerimeter(length, breadth float64) float64 {
	return 2. * (length + breadth)
}

func SumAreaPeri(length, breadth float64) float64 {
	return RectArea(length, breadth) + RectPerimeter(length, breadth)

}
