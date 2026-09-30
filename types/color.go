package types

import "regexp"

// Color patterns match go-playground/validator v10.28.0.
// iscolor there is the alias hexcolor|rgb|rgba|hsl|hsla.
// The sample comparison lives in opencodeco/validgen-benchmarks/color.
const (
	hexColorRegexString = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	rgbRegexString      = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegexString     = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegexString      = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegexString     = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
)

var (
	hexColorRegex = regexp.MustCompile(hexColorRegexString)
	rgbRegex      = regexp.MustCompile(rgbRegexString)
	rgbaRegex     = regexp.MustCompile(rgbaRegexString)
	hslRegex      = regexp.MustCompile(hslRegexString)
	hslaRegex     = regexp.MustCompile(hslaRegexString)
)

func IsHexColor(s string) bool {
	return hexColorRegex.MatchString(s)
}

func IsRGB(s string) bool {
	return rgbRegex.MatchString(s)
}

func IsRGBA(s string) bool {
	return rgbaRegex.MatchString(s)
}

func IsHSL(s string) bool {
	return hslRegex.MatchString(s)
}

func IsHSLA(s string) bool {
	return hslaRegex.MatchString(s)
}

// IsColor reports whether s matches hexcolor, rgb, rgba, hsl, or hsla.
func IsColor(s string) bool {
	return IsHexColor(s) || IsRGB(s) || IsRGBA(s) || IsHSL(s) || IsHSLA(s)
}
