package color

import "fmt"

type HSV struct {
	H, S, V uint32
}

// StrToHsv converts hsv() format string to HSV object
// Parameters:
//   str: string in "hsv(h,s,v)" format
// Returns:
//   *HSV: pointer to HSV object
//   error: parsing error if format is invalid
// Example:
//   c, err := StrToHsv("hsv(0,100,100)") // red
func StrToHsv(str string) (*HSV, error) {
	var h, s, v uint32
	_, err := fmt.Sscanf(RemoveSpace(str), "hsv(%d,%d,%d)", &h, &s, &v)
	if err != nil {
		return nil, err
	}
	return &HSV{H: h, S: s, V: v}, nil
}

// String converts HSV object to hsv() format string
// Returns:
//   string: "hsv(h,s,v)" formatted string
// Example:
//   c := HSV{240,100,100}
//   fmt.Println(c.String()) // outputs "hsv(240,100,100)"
func (c *HSV) String() string {
	return fmt.Sprintf("hsv(%d,%d,%d)", c.H, c.S, c.V)
}

// ToRgb converts HSV to RGB representation
// Returns:
//   RGB: corresponding RGB color object
// Example:
//   c := HSV{0,0,100} // white
//   rgb := c.ToRgb() // returns RGB{255,255,255}
func (c *HSV) ToRgb() RGB {
	r, g, b := hsvToRgb(c.H, c.S, c.V)
	return RGB{r, g, b}
}

// ToRgba converts HSV to RGBA with full opacity
// Returns:
//   RGBA: RGBA object with alpha=1.0
// Example:
//   c := HSV{120,100,100} // green
//   rgba := c.ToRgba() // returns RGBA{RGB{0,255,0},1.0}
func (c *HSV) ToRgba() RGBA {
	rgb := c.ToRgb()
	return RGBA{rgb, 1.0}
}

// ToHex converts HSV to hexadecimal string
// Returns:
//   string: "#RRGGBB" format string
// Example:
//   c := HSV{300,100,100} // magenta
//   hex := c.ToHex() // returns "#FF00FF"
func (c *HSV) ToHex() string {
	rgb := c.ToRgb()
	return rgb.ToHex()
}

// ToHsl converts HSV to HSL representation
// Returns:
//   Hsl: corresponding HSL color object
// Example:
//   c := HSV{60,100,100} // yellow
//   hsl := c.ToHsl() // returns Hsl{60,100,50}
func (c *HSV) ToHsl() HSL {
	rgb := c.ToRgb()
	return rgb.ToHsl()
}

// ToHsla converts HSV to HSLA with full opacity
// Returns:
//   HSLA: HSLA object with alpha=1.0
// Example:
//   c := HSV{0,100,50} // dark red
//   hsla := c.ToHsla() // returns HSLA{Hsl{0,100,25},1.0}
func (c *HSV) ToHsla() HSLA {
	rgb := c.ToRgb()
	return rgb.ToHsla()
}

// ToCmyk converts HSV to CMYK representation
// Returns:
//   CMYK: corresponding CMYK color object
// Example:
//   c := HSV{0,0,0} // black
//   cmyk := c.ToCmyk() // returns CMYK{0,0,0,100}
func (c *HSV) ToCmyk() CMYK {
	rgb := c.ToRgb()
	return rgb.ToCmyk()
}