package RectangleTester

import (
	"fmt"
	"testing"
)

func TestSquareArea(t *testing.T) {

	t.Run("Basic Tests for small  hard coded values", func(t *testing.T) {

		want := 9.
		got := Square{
			side: 3,
		}.Area()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
		}
	})
}

func TestSquarePerimeter(t *testing.T) {
	t.Run("Expect:12 Side:3.", func(t *testing.T) {
		want := 12.
		got := Square{
			side: 3.}.Perimeter()
		if want != got {
			t.Fail()
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
		}
	})
}

func TestSquare_SumAreaPerimeter(t *testing.T) {

	t.Run("Expect:21 Side:3.", func(t *testing.T) {
		want := 21.
		got := Square{
			side: 3.}.SumAreaPerimeter()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
		}
	})
}
