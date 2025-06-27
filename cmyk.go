package color

import "fmt"

type CMYK struct {
	C, M, Y, K uint8
}

// StrToCmyk converts a cmyk() format string to CMYK object
// Parameters:
//   str: string in "cmyk(c,m,y,k)" format
// Returns:
//   *CMYK: pointer to CMYK object
//   error: parsing error if format is invalid
// Example:
//   c, err := StrToCmyk("cmyk(100,0,0,0)") // pure cyan
func StrToCmyk(str string) (*CMYK, error) {
	var c, m, y, k uint8
	_, err := fmt.Sscanf(RemoveSpace(str), "cmyk(%d%%,%d%%,%d%%,%d%%)", &c, &m, &y, &k)
	if err != nil {
		return nil, err
	}
	return &CMYK{C: c, M: m, Y: y, K: k}, nil
}

// String converts CMYK object to cmyk() format string
// Returns:
//   string: "cmyk(c,m,y,k)" formatted string
// Example:
//   c := CMYK{0, 100, 100, 0}
//   fmt.Println(c.String()) // outputs "cmyk(0,100,100,0)"
func (c *CMYK) String() string {
	return fmt.Sprintf("cmyk(%d%%,%d%%,%d%%,%d%%)", c.C, c.M, c.Y, c.K)
}

// ToRgb converts CMYK to RGB representation
// Returns:
//   RGB: corresponding RGB color object
// Example:
//   c := CMYK{0, 100, 0, 0} // magenta
//   rgb := c.ToRgb() // returns RGB{255,0,255}
func (c *CMYK) ToRgb() RGB {
	r, g, b := cmykToRgb(c.C, c.M, c.Y, c.K)
	return RGB{r, g, b}
}

// ToRgba converts CMYK to RGBA with full opacity
// Returns:
//   RGBA: RGBA object with alpha=1.0
// Example:
//   c := CMYK{100,0,100,0} // green
//   rgba := c.ToRgba() // returns RGBA{RGB{0,255,0},1.0}
func (c *CMYK) ToRgba() RGBA {
	rgb := c.ToRgb()
	return RGBA{rgb, 1.0}
}

// ToHex converts CMYK to hexadecimal string
// Returns:
//   string: "#RRGGBB" format string
// Example:
//   c := CMYK{0,0,100,0} // yellow
//   hex := c.ToHex() // returns "#FFFF00"
func (c *CMYK) ToHex() string {
	rgb := c.ToRgb()
	return rgb.ToHex()
}

// ToHsl converts CMYK to HSL representation
// Returns:
//   Hsl: corresponding HSL color object
// Example:
//   c := CMYK{100,100,0,0} // blue
//   hsl := c.ToHsl() // returns Hsl{240,100,50}
func (c *CMYK) ToHsl() HSL {
	rgb := c.ToRgb()
	return rgb.ToHsl()
}

// ToHsla converts CMYK to HSLA with full opacity
// Returns:
//   HSLA: HSLA object with alpha=1.0
// Example:
//   c := CMYK{0,50,100,0} // orange
//   hsla := c.ToHsla() // returns HSLA{Hsl{30,100,50},1.0}
func (c *CMYK) ToHsla() HSLA {
	rgb := c.ToRgb()
	return rgb.ToHsla()
}

// ToHsv converts CMYK to HSV representation
// Returns:
//   HSV: corresponding HSV color object
// Example:
//   c := CMYK{0,0,0,100} // black
//   hsv := c.ToHsv() // returns HSV{0,0,0}
func (c *CMYK) ToHsv() HSV {
	rgb := c.ToRgb()
	return rgb.ToHsv()
}