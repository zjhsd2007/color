package color

import (
	"fmt"
)

type HSLA struct {
	HSL
	A float32
}

// StrToHsla converts hsla() format string to HSLA object
// Parameters:
//   str: string in "hsla(h,s%,l%,a)" format
// Returns:
//   *HSLA: pointer to HSLA object
//   error: parsing error if format is invalid
// Example:
//   c, err := StrToHsla("hsla(120,100%,50%,0.5)") // semi-transparent green
func StrToHsla(str string) (*HSLA, error) {
	var h, s, l uint32
	var a float32
	_, err := fmt.Sscanf(RemoveSpace(str), "hsla(%d,%d%%,%d%%,%f)", &h, &s, &l, &a)
	if err != nil {
		return nil, err
	}
	return &HSLA{HSL{h, s, l}, a}, nil
}

// String converts HSLA object to hsla() format string
// Returns:
//   string: "hsla(h,s%,l%,a)" formatted string
// Example:
//   c := HSLA{Hsl{0,100,50}, 0.75}
//   fmt.Println(c.String()) // outputs "hsla(0, 100%, 50%, 0.75)"
func (c *HSLA) String() string {
	return fmt.Sprintf("hsla(%d, %d%%, %d%%, %.2f)", c.H, c.S, c.L, c.A)
}

// ToRgb converts HSLA to RGB with alpha precomputation
// Returns:
//   RGB: RGB object with alpha applied
// Example:
//   c := HSLA{Hsl{240,100,50}, 0.8} // 80% opaque blue
//   rgb := c.ToRgb() // returns RGB{0,0,204} (with alpha calculation)
func (c *HSLA) ToRgb() RGB {
	r, g, b := hslToRgb(c.H, c.S, c.L)
	return RGB{
		calcRgbWithAlpha(r, c.A),
		calcRgbWithAlpha(g, c.A),
		calcRgbWithAlpha(b, c.A),
	}
}

// ToRgba converts HSLA to RGBA representation
// Returns:
//   RGBA: RGBA object with preserved alpha
// Example:
//   c := HSLA{Hsl{60,100,50}, 0.6} // 60% opaque yellow
//   rgba := c.ToRgba() // returns RGBA{RGB{255,255,0},0.6}
func (c *HSLA) ToRgba() RGBA {
	rgb := c.ToRgb()
	return RGBA{rgb, c.A}
}

// ToHex converts HSLA to hexadecimal string (alpha not included)
// Returns:
//   string: "#RRGGBB" format string
// Example:
//   c := HSLA{Hsl{0,100,50}, 1.0} // red
//   hex := c.ToHex() // returns "#FF0000"
func (c *HSLA) ToHex() string {
	rgb := c.ToRgb()
	return rgb.ToHex()
}

// ToHsl converts HSLA to HSL (discards alpha)
// Returns:
//   Hsl: HSL object without alpha
// Example:
//   c := HSLA{Hsl{180,100,50}, 0.5}
//   hsl := c.ToHsl() // returns Hsl{180,100,50}
func (c *HSLA) ToHsl() HSL {
	rgb := c.ToRgb()
	return rgb.ToHsl()
}

// ToHsv converts HSLA to HSV representation
// Returns:
//   HSV: corresponding HSV color object
// Example:
//   c := HSLA{Hsl{300,100,50}, 1.0} // magenta
//   hsv := c.ToHsv() // returns HSV{300,100,100}
func (c *HSLA) ToHsv() HSV {
	rgb := c.ToRgb()
	return rgb.ToHsv()
}

// ToCmyk converts HSLA to CMYK representation
// Returns:
//   CMYK: corresponding CMYK color object
// Example:
//   c := HSLA{Hsl{120,100,25}, 1.0} // dark green
//   cmyk := c.ToCmyk() // returns CMYK{100,0,100,50}
func (c *HSLA) ToCmyk() CMYK {
	rgb := c.ToRgb()
	return rgb.ToCmyk()
}