package color

type HEX struct {
	str string
	rgb [3]uint8
	a float32
}

// StrToHex converts a hexadecimal color string to a Hex object
// Parameters:
//   str: string in "#RRGGBB" or "#RRGGBBAA" format
// Returns:
//   *Hex: pointer to converted Hex object
//   error: error information if conversion fails
// Example:
//   hex, err := StrToHex("#FF5733") // returns Hex object for orange
//   hex, err := StrToHex("#FF573380") // orange with transparency
func StrToHex(str string) (*HEX, error) {
	r, g, b, a, err := hexToRGBA(str)
	if err != nil {
		return nil, err
	}
	return &HEX{str, [3]uint8{r,g,b}, float32(a/255)}, nil
}

// String converts Hex object to hexadecimal string representation
// Returns:
//   string: "#RRGGBB" format string (transparency ignored)
// Example:
//   c := Hex{RGBA{RGB{255, 87, 51}, 1.0}}
//   fmt.Println(c.String()) // outputs "#FF5733"
func (c *HEX) String() string {
	return c.str
}

// ToRgb converts Hex object to RGB object (with alpha calculation applied)
// Returns:
//   RGB: RGB object after alpha calculation
// Example:
//   c := Hex{RGBA{RGB{255, 0, 0}, 0.5}} // semi-transparent red
//   rgb := c.ToRgb() // returns RGB{128, 0, 0}
func (c *HEX) ToRgb() RGB {
	return RGB{
		calcRgbWithAlpha(c.rgb[0], c.a),
		calcRgbWithAlpha(c.rgb[1], c.a),
		calcRgbWithAlpha(c.rgb[2], c.a),
	}
}

// ToRgba converts Hex object to RGBA object (opacity fixed at 1.0)
// Returns:
//   RGBA: RGBA object with fixed 1.0 opacity
// Example:
//   c := Hex{RGBA{RGB{0, 255, 0}, 0.5}} // semi-transparent green
//   rgba := c.ToRgba() // returns RGBA{RGB{0,255,0}, 1.0}
func (c *HEX) ToRgba() RGBA {
	return RGBA{c.ToRgb(), 1.0}
}

// ToHsl converts Hex object to Hsl object (hue-saturation-lightness)
// Returns:
//   Hsl: corresponding HSL color object
// Example:
//   c := Hex{RGBA{RGB{255, 0, 0}, 1.0}} // red
//   hsl := c.ToHsl() // returns Hsl{0, 100, 50}
func (c *HEX) ToHsl() HSL {
	rgb := c.ToRgb()
	return rgb.ToHsl()
}

// ToHsla converts Hex object to HSLA object (with transparency)
// Returns:
//   HSLA: corresponding HSLA color object
// Example:
//   c := Hex{RGBA{RGB{0, 0, 255}, 0.8}} // 80% opaque blue
//   hsla := c.ToHsla() // returns HSLA{Hsl{240,100,50}, 0.8}
func (c *HEX) ToHsla() HSLA {
	rgba := c.ToRgba()
	return rgba.ToHsla()
}

// ToHsv converts Hex object to HSV object (hue-saturation-value)
// Returns:
//   HSV: corresponding HSV color object
// Example:
//   c := Hex{RGBA{RGB{255, 255, 0}, 1.0}} // yellow
//   hsv := c.ToHsv() // returns HSV{60, 100, 100}
func (c *HEX) ToHsv() HSV {
	rgb := c.ToRgb()
	return rgb.ToHsv()
}

// ToCmyk converts Hex object to CMYK object (cyan-magenta-yellow-key)
// Returns:
//   CMYK: corresponding CMYK color object
// Example:
//   c := Hex{RGBA{RGB{255, 0, 255}, 1.0}} // magenta
//   cmyk := c.ToCmyk() // returns CMYK{0, 100, 0, 0}
func (c *HEX) ToCmyk() CMYK {
	rgb := c.ToRgb()
	return rgb.ToCmyk()
}