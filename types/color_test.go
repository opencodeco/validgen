package types

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestColorMatchesValidator(t *testing.T) {
	validate := validator.New()
	samples := []string{
		"",
		"#000",
		"#000-",
		"#fff",
		"#FFFF",
		"#aabbcc",
		"#aabbccdd",
		"#ab",
		"#abcde",
		"000000",
		"rgb(0,0,0)",
		"rgb(255,255,255)",
		"rgb(255, 255, 255)",
		"rgb(0%,0%,0%)",
		"rgb(100%, 0%, 50%)",
		"rgb(255%,255%,255%)",
		"rgb(256,0,0)",
		"rgb(0, 0, 0, 1)",
		"RGB(0,0,0)",
		"rgba(0,0,0,0)",
		"rgba(0,0,0,1)",
		"rgba(255,255,255,0.5)",
		"rgba(0,0,0,0.0)",
		"rgba(0%,0%,0%,1)",
		"rgba(0,0,0,2)",
		"hsl(0,0%,0%)",
		"hsl(360,100%,50%)",
		"hsl(0, 0%, 0%)",
		"hsl(361,0%,0%)",
		"hsl(120,50,50)",
		"hsla(120,100%,50%,1)",
		"hsla(120,100%,50%,0.5)",
		"hsla(120,100%,50%,0.0)",
		"red",
		"blue",
		"cmyk(0%,0%,0%,0%)",
		"rgb(0,0,0%)",
	}
	checks := []struct {
		tag string
		fn  func(string) bool
	}{
		{tag: "hexcolor", fn: IsHexColor},
		{tag: "rgb", fn: IsRGB},
		{tag: "rgba", fn: IsRGBA},
		{tag: "hsl", fn: IsHSL},
		{tag: "hsla", fn: IsHSLA},
		{tag: "iscolor", fn: IsColor},
	}

	for _, sample := range samples {
		for _, check := range checks {
			err := validate.Var(sample, check.tag)
			got := check.fn(sample)
			want := err == nil
			if got != want {
				t.Errorf("%s %q = %v, validator accepted %v", check.tag, sample, got, want)
			}
		}
	}
}

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
