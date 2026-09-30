package types

import "testing"

func TestCommonExampleColors(t *testing.T) {
	if IsColor("#000-") {
		t.Fatal(`#000- is the invalid color in the common validator example`)
	}
	if !IsColor("#000") {
		t.Fatal(`#000 should be a color`)
	}
	if !IsColor("rgb(255,255,255)") {
		t.Fatal(`rgb(255,255,255) should be a color`)
	}
	if IsColor("red") {
		t.Fatal(`named colors are outside hexcolor|rgb|rgba|hsl|hsla`)
	}
}
