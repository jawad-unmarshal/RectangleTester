package RectangleTester

import (
	"fmt"
	"testing"
)

func TestRectArea(t *testing.T) {

	t.Run("Basic Tests for small  hard coded values", func(t *testing.T) {

		want := 12.
		got := Rectangle{
			length:  3,
			breadth: 4,
		}.Area()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
		}
	})

	t.Run("Testing for non hard-coded values", func(t *testing.T) {
		want := 144.
		got := Rectangle{
			length:  12,
			breadth: 12,
		}.Area()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
		}
	})
}

func TestRectPerimeter(t *testing.T) {
	t.Run("Expect:14 Len:3. Breadth:4.", func(t *testing.T) {
		want := 14.
		got := Rectangle{
			length:  3.,
			breadth: 4.}.Perimeter()
		if want != got {
			t.Fail()
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
		}
	})
	t.Run("Expect:64 Len:20. Breadth:12.", func(t *testing.T) {
		want := 64.
		got := Rectangle{
			length:  20.,
			breadth: 12.}.Perimeter()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
			//t.Fail()
		}
	})
}

func TestSumAreaPerimeter(t *testing.T) {

	t.Run("Expect:26 Len:3. Breadth:4.", func(t *testing.T) {
		want := 26.
		got := Rectangle{
			length:  3.,
			breadth: 4.}.SumAreaPerimeter()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
			//t.Fail()Fail

		}
	})

	t.Run("Expect:304 Len:20. Breadth:12.", func(t *testing.T) {
		want := 304.
		got := Rectangle{
			length:  20.,
			breadth: 12.}.SumAreaPerimeter()
		if want != got {
			t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
			//t.Fail()
		}
	})

}
