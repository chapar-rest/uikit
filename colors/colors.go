package colors

import (
	"fmt"
	"image/color"
	"strings"
)

var (
	Gray        = color.NRGBA{R: 120, G: 120, B: 120, A: 255}
	Black       = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	White       = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	Red         = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	Green       = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	Blue        = color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	Yellow      = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	Pink        = color.NRGBA{R: 255, G: 192, B: 203, A: 255}
	Purple      = color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	Orange      = color.NRGBA{R: 255, G: 165, B: 0, A: 255}
	Brown       = color.NRGBA{R: 165, G: 42, B: 42, A: 255}
	LightGray   = color.NRGBA{R: 211, G: 211, B: 211, A: 255}
	DarkGray    = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	LightBlue   = color.NRGBA{R: 173, G: 216, B: 230, A: 255}
	LightGreen  = color.NRGBA{R: 144, G: 238, B: 144, A: 255}
	LightRed    = color.NRGBA{R: 255, G: 105, B: 97, A: 255}
	LightYellow = color.NRGBA{R: 255, G: 255, B: 224, A: 255}
	LightPink   = color.NRGBA{R: 255, G: 192, B: 203, A: 255}
	LightPurple = color.NRGBA{R: 173, G: 139, B: 168, A: 255}
	LightOrange = color.NRGBA{R: 255, G: 165, B: 0, A: 255}
	LightBrown  = color.NRGBA{R: 165, G: 42, B: 42, A: 255}
)

func FromHex(hex string) color.NRGBA {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		panic("hex must be 6 characters")
	}

	r, g, b, a := uint8(0), uint8(0), uint8(0), uint8(255)
	if _, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b); err != nil {
		panic(err)
	}
	return color.NRGBA{R: r, G: g, B: b, A: a}
}
