package color

import "fmt"

type HSL struct {
	H, S, L uint32
}

// StrToHsl converts hsl() format string to Hsl object
// Parameters:
//   str: string in "hsl(h,s,l)" format
// Returns:
//   *Hsl: pointer to Hsl object
//   error: parsing error if format is invalid
// Example:
//   c, err := StrToHsl("hsl(0,100,50)") // red
func StrToHsl(str string) (*HSL, error) {
	var h, s, l uint32
	_, err := fmt.Sscanf(RemoveSpace(str), "hsl(%d,%d%%,%d%%)", &h, &s, &l)
	if err != nil {
		return nil, err
	}
	return &HSL{H: h, S: s, L: l}, nil
}

// String converts Hsl object to hsl() format string
// Returns:
//   string: "hsl(h,s,l)" formatted string
// Example:
//   c := Hsl{120,100,50}
//   fmt.Println(c.String()) // outputs "hsl(120,100%,50%)"
func (c *HSL) String() string {
	return fmt.Sprintf("hsl(%d, %d%%, %d%%)", c.H, c.S, c.L)
}

// ToRgb converts HSL to RGB representation
// Returns:
//   RGB: corresponding RGB color object
// Example:
//   c := Hsl{0,0,50} // medium gray
//   rgb := c.ToRgb() // returns RGB{128,128,128}
func (c *HSL) ToRgb() RGB {
	r, g, b := hslToRgb(c.H, c.S, c.L)
	return RGB{R: r, G: g, B: b}
}

// ToRgba converts HSL to RGBA with full opacity
// Returns:
//   RGBA: RGBA object with alpha=1.0
// Example:
//   c := Hsl{30,100,50} // orange
//   rgba := c.ToRgba() // returns RGBA{RGB{255,128,0},1.0}
func (c *HSL) ToRgba() RGBA {
	rgb := c.ToRgb()
	return RGBA{rgb, 1.0}
}

// ToHex converts HSL to hexadecimal string
// Returns:
//   string: "#RRGGBB" format string
// Example:
//   c := Hsl{270,100,40} // purple
//   hex := c.ToHex() // returns "#6600CC"
func (c *HSL) ToHex() string {
	rgb := c.ToRgb()
	return rgb.ToHex()
}

// ToHsla converts HSL to HSLA with full opacity
// Returns:
//   HSLA: HSLA object with alpha=1.0
// Example:
//   c := Hsl{0,100,25} // dark red
//   hsla := c.ToHsla() // returns HSLA{Hsl{0,100,25},1.0}
func (c *HSL) ToHsla() HSLA {
	rgba := c.ToRgba()
	return rgba.ToHsla()
}

// ToHsv converts HSL to HSV representation
// Returns:
//   HSV: corresponding HSV color object
// Example:
//   c := Hsl{180,100,50} // cyan
//   hsv := c.ToHsv() // returns HSV{180,100,100}
func (c *HSL) ToHsv() HSV {
	rgb := c.ToRgb()
	return rgb.ToHsv()
}

// ToCmyk converts HSL to CMYK representation
// Returns:
//   CMYK: corresponding CMYK color object
// Example:
//   c := Hsl{0,0,0} // black
//   cmyk := c.ToCmyk() // returns CMYK{0,0,0,100}
func (c *HSL) ToCmyk() CMYK {
	rgb := c.ToRgb()
	return rgb.ToCmyk()
}