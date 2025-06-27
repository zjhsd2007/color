package color

import "fmt"

type RGBA struct {
	RGB
	A float32
}

// StrToRgba converts an rgba() format string to RGBA object
// Parameters:
//   str: string in "rgba(r,g,b,a)" format
// Returns:
//   *RGBA: pointer to converted RGBA object
//   error: parsing error if format is invalid
// Example:
//   c, err := StrToRgba("rgba(255,87,51,0.8)") // orange with 80% opacity
func StrToRgba(str string) (*RGBA, error) {
	var r, g, b uint8
	var a float32
	_, err := fmt.Sscanf(RemoveSpace(str), "rgba(%d,%d,%d,%f)", &r, &g, &b, &a)
	if err != nil {
		return nil, err
	}
	return &RGBA{RGB{r, g, b}, a}, nil
}

// String converts RGBA object to rgba() format string
// Returns:
//   string: "rgba(r,g,b,a)" formatted string
// Example:
//   c := RGBA{RGB{0, 128, 255}, 0.6}
//   fmt.Println(c.String()) // outputs "rgba(0,128,255,0.600000)"
func (c *RGBA) String() string {
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", c.R, c.G, c.B, c.A)
}

// ToRgb converts RGBA to RGB with alpha precomputation
// Returns:
//   RGB: RGB object with alpha applied to channels
// Example:
//   c := RGBA{RGB{255, 0, 0}, 0.5} // semi-transparent red
//   rgb := c.ToRgb() // returns RGB{128,0,0}
func (c *RGBA) ToRgb() RGB {
	return RGB{
		R: calcRgbWithAlpha(c.R, c.A),
		G: calcRgbWithAlpha(c.G, c.A),
		B: calcRgbWithAlpha(c.B, c.A),
	}
}

// ToHex converts RGBA object to hexadecimal string
// Returns:
//   string: "#RRGGBB" format (alpha not included)
// Example:
//   c := RGBA{RGB{0, 255, 127}, 1.0}
//   hex := c.ToHex() // returns "#00FF7F"
func (c *RGBA) ToHex() string {
	rgb := c.ToRgb()
	return rgb.ToHex()
}

// ToHsl converts RGBA to HSL representation
// Returns:
//   Hsl: HSL color object (alpha not included)
// Example:
//   c := RGBA{RGB{0, 0, 255}, 1.0} // blue
//   hsl := c.ToHsl() // returns Hsl{240,100,50}
func (c *RGBA) ToHsl() HSL {
	rgb := c.ToRgb()
	return rgb.ToHsl()
}

// ToHsla converts RGBA to HSLA representation
// Returns:
//   HSLA: HSLA object with preserved alpha
// Example:
//   c := RGBA{RGB{255, 165, 0}, 0.7} // orange with 70% opacity
//   hsla := c.ToHsla() // returns HSLA{Hsl{38.8,100,50}, 0.7}
func (c *RGBA) ToHsla() HSLA {
	h, s, l, a := rgbaToHsla(c.R, c.G, c.B, c.A)
	return HSLA{HSL{h, s, l}, a}
}

// ToHsv converts RGBA to HSV representation
// Returns:
//   HSV: HSV color object
// Example:
//   c := RGBA{RGB{255, 0, 255}, 1.0} // magenta
//   hsv := c.ToHsv() // returns HSV{300,100,100}
func (c *RGBA) ToHsv() HSV {
	rgb := c.ToRgb()
	return rgb.ToHsv()
}

// ToCmyk converts RGBA to CMYK representation
// Returns:
//   CMYK: CMYK color object
// Example:
//   c := RGBA{RGB{0, 255, 255}, 1.0} // cyan
//   cmyk := c.ToCmyk() // returns CMYK{100,0,0,0}
func (c *RGBA) ToCmyk() CMYK {
	r1 := calcRgbWithAlpha(c.R, c.A)
	g1 := calcRgbWithAlpha(c.G, c.A)
	b1 := calcRgbWithAlpha(c.B, c.A)
	cv, m, y, k := rgbToCmyk(r1, g1, b1)
	return CMYK{cv, m, y, k}
}