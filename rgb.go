package color

import "fmt"

type RGB struct {
	R uint8
	G uint8
	B uint8
}

// StrToRgb converts rgb() format string to RGB object
// Parameters:
//   str: string in "rgb(r,g,b)" format
// Returns:
//   *RGB: pointer to RGB object
//   error: parsing error if format is invalid
// Example:
//   c, err := StrToRgb("rgb(0,0,255)") // blue
func StrToRgb(str string) (*RGB, error) {
	var r, g, b uint8
	_, err := fmt.Sscanf(RemoveSpace(str), "rgb(%d,%d,%d)", &r, &g, &b)
	if err != nil {
		return nil, err
	}
	return &RGB{R: r, G: g, B: b}, nil
}

// String converts RGB object to rgb() format string
// Returns:
//   string: "rgb(r,g,b)" formatted string
// Example:
//   c := RGB{192,192,192}
//   fmt.Println(c.String()) // outputs "rgb(192,192,192)"
func (c *RGB) String() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

// ToRgba converts RGB to RGBA with full opacity
// Returns:
//   RGBA: RGBA object with alpha=1.0
// Example:
//   c := RGB{255,165,0} // orange
//   rgba := c.ToRgba() // returns RGBA{RGB{255,165,0},1.0}
func (c *RGB) ToRgba() RGBA {
	return RGBA{*c, 1.0}
}

// ToHex converts RGB to hexadecimal string
// Returns:
//   string: "#RRGGBB" format string
// Example:
//   c := RGB{128,0,128} // purple
//   hex := c.ToHex() // returns "#800080"
func (c *RGB) ToHex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// ToHsl converts RGB to HSL representation
// Returns:
//   Hsl: corresponding HSL color object
// Example:
//   c := RGB{255,0,0} // red
//   hsl := c.ToHsl() // returns Hsl{0,100,50}
func (c *RGB) ToHsl() HSL {
	h, s, l := rgbToHsl(c.R, c.G, c.B)
	return HSL{H: h, S: s, L: l}
}

// ToHsla converts RGB to HSLA with full opacity
// Returns:
//   HSLA: HSLA object with alpha=1.0
// Example:
//   c := RGB{0,255,0} // green
//   hsla := c.ToHsla() // returns HSLA{Hsl{120,100,50},1.0}
func (c *RGB) ToHsla() HSLA {
	return HSLA{
		c.ToHsl(),
		1.0,
	}
}

// ToHsv converts RGB to HSV representation
// Returns:
//   HSV: corresponding HSV color object
// Example:
//   c := RGB{0,0,255} // blue
//   hsv := c.ToHsv() // returns HSV{240,100,100}
func (c *RGB) ToHsv() HSV {
	h, s, v := rgbToHsv(c.R, c.G, c.B)
	return HSV{H: h, S: s, V: v}
}

// ToCmyk converts RGB to CMYK representation
// Returns:
//   CMYK: corresponding CMYK color object
// Example:
//   c := RGB{255,255,0} // yellow
//   cmyk := c.ToCmyk() // returns CMYK{0,0,100,0}
func (c *RGB) ToCmyk() CMYK {
	cv, m, y, k := rgbToCmyk(c.R, c.G, c.B)
	return CMYK{C: cv, M: m, Y: y, K: k}
}