package RectangleTester

import (
	"fmt"
	"testing"
)

func TestRectArea(t *testing.T) {
	//name:= "Basic Tests for small  hard coded values"
	want := 12.
	got := RectArea(3., 4.)
	if want != got {
		t.Fail()
		t.Errorf(fmt.Sprintf("Wanted %v\t Got%v", want, got))
	}

}
