package color

import (
	"testing"
)

func TestRgb(t *testing.T) {
	t.Run("string to rgb", func(t *testing.T) {
		rgb, _ := StrToRgb("rgb(255,255,255)")
		expected := RGB{255, 255, 255}
		if *rgb != expected {
			t.Errorf("expected %v, got %v", expected, *rgb)
		}
	})

	t.Run("rgb to string", func(t *testing.T) {
		c := RGB{0, 128, 255}
		str := c.String()
		expected := "rgb(0,128,255)"
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("rgb to rgba", func(t *testing.T) {
		c := RGB{192, 192, 192}
		rgba := c.ToRgba()
		expected := RGBA{RGB{192, 192, 192}, 1.0}
		if rgba != expected {
			t.Errorf("expected %v, got %v", expected, rgba)
		}
	})

	t.Run("rgb to hex", func(t *testing.T) {
		c := RGB{255, 0, 127}
		hex := c.ToHex()
		expected := "#ff007f"
		if hex != expected {
			t.Errorf("expected %s, got %s", expected, hex)
		}
	})

	t.Run("rgb to hsl", func(t *testing.T) {
		c := RGB{255, 0, 0} // red
		hsl := c.ToHsl()
		expected := HSL{0, 100, 50}
		if hsl != expected {
			t.Errorf("expected %v, got %v", expected, hsl)
		}
	})

	t.Run("rgb to hsla", func(t *testing.T) {
		c := RGB{0, 255, 0} // green
		hsla := c.ToHsla()
		expected := HSLA{HSL{120, 100, 50}, 1.0}
		if hsla != expected {
			t.Errorf("expected %v, got %v", expected, hsla)
		}
	})

	t.Run("rgb to hsv", func(t *testing.T) {
		c := RGB{0, 0, 255} // blue
		hsv := c.ToHsv()
		expected := HSV{240, 100, 100}
		if hsv != expected {
			t.Errorf("expected %v, got %v", expected, hsv)
		}
	})

	t.Run("rgb to cmyk", func(t *testing.T) {
		c := RGB{255, 255, 0} // yellow
		cmyk := c.ToCmyk()
		expected := CMYK{0, 0, 100, 0}
		if cmyk != expected {
			t.Errorf("expected %v, got %v", expected, cmyk)
		}
	})
}

func TestRgba(t *testing.T) {
	t.Run("string to rgba", func(t *testing.T) {
		rgba, _ := StrToRgba("rgba(255,87,51,0.8)")
		expected := RGBA{RGB{255, 87, 51}, 0.8}
		if rgba.R != expected.R || rgba.G != expected.G || rgba.B != expected.B || rgba.A != expected.A {
			t.Errorf("expected %v, got %v", expected, *rgba)
		}
	})

	t.Run("rgba to string", func(t *testing.T) {
		c := RGBA{RGB{0, 128, 255}, 0.6}
		str := c.String()
		expected := "rgba(0,128,255,0.60)"
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("rgba to rgb", func(t *testing.T) {
		c := RGBA{RGB{255, 0, 0}, 0.5} // semi-transparent red
		rgb := c.ToRgb()
		expected := RGB{255, 128, 128}
		if rgb != expected {
			t.Errorf("expected %v, got %v", expected, rgb)
		}
	})

	t.Run("rgba to hex", func(t *testing.T) {
		c := RGBA{RGB{0, 255, 127}, 1.0}
		hex := c.ToHex()
		expected := "#00ff7f"
		if hex != expected {
			t.Errorf("expected %s, got %s", expected, hex)
		}
	})

	t.Run("rgba to hsl", func(t *testing.T) {
		c := RGBA{RGB{0, 0, 255}, 1.0} // blue
		hsl := c.ToHsl()
		expected := HSL{240, 100, 50}
		if hsl != expected {
			t.Errorf("expected %v, got %v", expected, hsl)
		}
	})

	t.Run("rgba to hsla", func(t *testing.T) {
		c := RGBA{RGB{255, 165, 0}, 0.7} // orange with 70% opacity
		hsla := c.ToHsla()
		expected := HSLA{HSL{38, 100, 50}, 0.7}
		if hsla.H != expected.H || hsla.A != expected.A {
			t.Errorf("expected %v, got %v", expected, hsla)
		}
	})

	t.Run("rgba to hsv", func(t *testing.T) {
		c := RGBA{RGB{255, 0, 255}, 1.0} // magenta
		hsv := c.ToHsv()
		expected := HSV{300, 100, 100}
		if hsv != expected {
			t.Errorf("expected %v, got %v", expected, hsv)
		}
	})

	t.Run("rgba to cmyk", func(t *testing.T) {
		c := RGBA{RGB{0, 255, 255}, 1.0} // cyan
		cmyk := c.ToCmyk()
		expected := CMYK{100, 0, 0, 0}
		if cmyk != expected {
			t.Errorf("expected %v, got %v", expected, cmyk)
		}
	})
}

func TestHex(t *testing.T) {
	t.Run("string to hex", func(t *testing.T) {
		hex, err := StrToHex("#ff007f")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := HEX{str: "#ff007f", rgb: [3]uint8{255, 0, 127}, a: 1.0}
		if hex.str != expected.str || hex.rgb != expected.rgb || hex.a != expected.a {
			t.Errorf("expected %+v, got %+v", expected, hex)
		}
	})

	t.Run("hex to string", func(t *testing.T) {
		h := HEX{str: "#00ff7f", rgb: [3]uint8{0, 255, 127}, a: 1.0}
		str := h.String()
		expected := "#00ff7f"
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("hex to rgb", func(t *testing.T) {
		h := HEX{str: "#ff0000", rgb: [3]uint8{255, 0, 0}, a: 1.0}
		rgb := h.ToRgb()
		expected := RGB{255, 0, 0}
		if rgb != expected {
			t.Errorf("expected %v, got %v", expected, rgb)
		}
	})

	t.Run("hex to rgba", func(t *testing.T) {
		h := HEX{str: "#8000ff80", rgb: [3]uint8{128, 0, 255}, a: 0.5}
		rgba := h.ToRgba()
		expected := RGBA{RGB{192, 128, 255}, 1.0}
		if rgba != expected {
			t.Errorf("expected %v, got %v", expected, rgba)
		}
	})

	t.Run("hex to hsl", func(t *testing.T) {
		h := HEX{str: "#00ffff", rgb: [3]uint8{0, 255, 255}, a: 1.0}
		hsl := h.ToHsl()
		expected := HSL{180, 100, 50}
		if hsl != expected {
			t.Errorf("expected %v, got %v", expected, hsl)
		}
	})

	t.Run("hex to hsla", func(t *testing.T) {
		h := HEX{str: "#ffa50080", rgb: [3]uint8{255, 165, 0}, a: 0.5}
		hsla := h.ToHsla()
		expected := HSLA{HSL{38, 99, 75}, 1.0}
		if hsla != expected {
			t.Errorf("expected %v, got %v", expected, hsla)
		}
	})

	t.Run("hex to hsv", func(t *testing.T) {
		h := HEX{str: "#800080", rgb: [3]uint8{128, 0, 128}, a: 1.0}
		hsv := h.ToHsv()
		expected := HSV{300, 100, 50}
		if hsv != expected {
			t.Errorf("expected %v, got %v", expected, hsv)
		}
	})

	t.Run("hex to cmyk", func(t *testing.T) {
		h := HEX{str: "#00ff00", rgb: [3]uint8{0, 255, 0}, a: 1.0}
		cmyk := h.ToCmyk()
		expected := CMYK{100, 0, 100, 0}
		if cmyk != expected {
			t.Errorf("expected %v, got %v", expected, cmyk)
		}
	})
}

func TestCmyk(t *testing.T) {
	t.Run("string to cmyk", func(t *testing.T) {
		cmyk, _ := StrToCmyk("cmyk(100%,0%,0%,0%)")
		expected := CMYK{100, 0, 0, 0}
		if *cmyk != expected {
			t.Errorf("expected %v, got %v", expected, *cmyk)
		}
	})

	t.Run("cmyk to string", func(t *testing.T) {
		c := CMYK{0, 100, 100, 0}
		str := c.String()
		expected := "cmyk(0%,100%,100%,0%)"
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("cmyk to rgb", func(t *testing.T) {
		c := CMYK{0, 100, 0, 0} // magenta
		rgb := c.ToRgb()
		expected := RGB{255, 0, 255}
		if rgb != expected {
			t.Errorf("expected %v, got %v", expected, rgb)
		}
	})

	t.Run("cmyk to rgba", func(t *testing.T) {
		c := CMYK{100, 0, 100, 0} // green
		rgba := c.ToRgba()
		expected := RGBA{RGB{0, 255, 0}, 1.0}
		if rgba.R != expected.R || rgba.G != expected.G || rgba.B != expected.B || rgba.A != expected.A {
			t.Errorf("expected %v, got %v", expected, rgba)
		}
	})

	t.Run("cmyk to hex", func(t *testing.T) {
		c := CMYK{0, 0, 100, 0} // yellow
		hex := c.ToHex()
		expected := "#ffff00"
		if hex != expected {
			t.Errorf("expected %s, got %s", expected, hex)
		}
	})

	t.Run("cmyk to hsl", func(t *testing.T) {
		c := CMYK{100, 100, 0, 0} // blue
		hsl := c.ToHsl()
		expected := HSL{240, 100, 50}
		if hsl != expected {
			t.Errorf("expected %v, got %v", expected, hsl)
		}
	})

	t.Run("cmyk to hsla", func(t *testing.T) {
		c := CMYK{0, 50, 100, 0} // orange
		hsla := c.ToHsla()
		expected := HSLA{HSL{30, 100, 50}, 1.0}
		if hsla.H != expected.H || hsla.S != expected.S || hsla.L != expected.L || hsla.A != expected.A {
			t.Errorf("expected %v, got %v", expected, hsla)
		}
	})

	t.Run("cmyk to hsv", func(t *testing.T) {
		c := CMYK{0, 0, 0, 100} // black
		hsv := c.ToHsv()
		expected := HSV{0, 0, 0}
		if hsv != expected {
			t.Errorf("expected %v, got %v", expected, hsv)
		}
	})
}

func TestHsl(t *testing.T) {
	t.Run("string to hsl", func(t *testing.T) {
		hsl, _ := StrToHsl("hsl(0,100%,50%)")
		expected := HSL{0, 100, 50}
		if *hsl != expected {
			t.Errorf("expected %v, got %v", expected, *hsl)
		}
	})

	t.Run("hsl to string", func(t *testing.T) {
		c := HSL{120, 100, 50}
		str := c.String()
		expected := "hsl(120, 100%, 50%)"
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("hsl to rgb", func(t *testing.T) {
		c := HSL{0, 0, 50} // medium gray
		rgb := c.ToRgb()
		expected := RGB{128, 128, 128}
		if rgb != expected {
			t.Errorf("expected %v, got %v", expected, rgb)
		}
	})

	t.Run("hsl to rgba", func(t *testing.T) {
		c := HSL{30, 100, 50} // orange
		rgba := c.ToRgba()
		expected := RGBA{RGB{255, 128, 0}, 1.0}
		if rgba.R != expected.R || rgba.G != expected.G || rgba.B != expected.B || rgba.A != expected.A {
			t.Errorf("expected %v, got %v", expected, rgba)
		}
	})

	t.Run("hsl to hex", func(t *testing.T) {
		c := HSL{270, 100, 40} // purple
		hex := c.ToHex()
		expected := "#6600cc"
		if hex != expected {
			t.Errorf("expected %s, got %s", expected, hex)
		}
	})

	t.Run("hsl to hsla", func(t *testing.T) {
		c := HSL{0, 100, 25} // dark red
		hsla := c.ToHsla()
		expected := HSLA{HSL{0, 100, 25}, 1.0}
		if hsla != expected {
			t.Errorf("expected %v, got %v", expected, hsla)
		}
	})

	t.Run("hsl to hsv", func(t *testing.T) {
		c := HSL{180, 100, 50} // cyan
		hsv := c.ToHsv()
		expected := HSV{180, 100, 100}
		if hsv != expected {
			t.Errorf("expected %v, got %v", expected, hsv)
		}
	})

	t.Run("hsl to cmyk", func(t *testing.T) {
		c := HSL{0, 0, 0} // black
		cmyk := c.ToCmyk()
		expected := CMYK{0, 0, 0, 100}
		if cmyk != expected {
			t.Errorf("expected %v, got %v", expected, cmyk)
		}
	})
}

func TestHsla(t *testing.T) {
	t.Run("string to hsla", func(t *testing.T) {
		hsla, _ := StrToHsla("hsla(120,100%,50%,0.5)")
		expected := HSLA{HSL{120, 100, 50}, 0.5}
		if hsla.H != expected.H || hsla.S != expected.S || hsla.L != expected.L || hsla.A != expected.A {
			t.Errorf("expected %v, got %v", expected, *hsla)
		}
	})

	t.Run("hsla to string", func(t *testing.T) {
		c := HSLA{HSL{0, 100, 50}, 0.75}
		str := c.String()
		expected := "hsla(0, 100%, 50%, 0.75)" // 注意格式匹配
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("hsla to rgb", func(t *testing.T) {
		c := HSLA{HSL{240, 100, 50}, 0.8} // 80% opaque blue
		rgb := c.ToRgb()
		expected := RGB{51, 51, 255}
		if rgb != expected {
			t.Errorf("expected %v, got %v", expected, rgb)
		}
	})

	t.Run("hsla to rgba", func(t *testing.T) {
		c := HSLA{HSL{60, 100, 50}, 0.6} // 60% opaque yellow
		rgba := c.ToRgba()
		expected := RGBA{RGB{255, 255, 102}, 0.6}
		if rgba.R != expected.R || rgba.G != expected.G || rgba.B != expected.B || rgba.A != expected.A {
			t.Errorf("expected %v, got %v", expected, rgba)
		}
	})

	t.Run("hsla to hex", func(t *testing.T) {
		c := HSLA{HSL{0, 100, 50}, 1.0} // red
		hex := c.ToHex()
		expected := "#ff0000"
		if hex != expected {
			t.Errorf("expected %s, got %s", expected, hex)
		}
	})

	t.Run("hsla to hsl", func(t *testing.T) {
		c := HSLA{HSL{180, 100, 50}, 0.5}
		hsl := c.ToHsl()
		expected := HSL{180, 99, 75}
		if hsl != expected {
			t.Errorf("expected %v, got %v", expected, hsl)
		}
	})

	t.Run("hsla to hsv", func(t *testing.T) {
		c := HSLA{HSL{300, 100, 50}, 1.0} // magenta
		hsv := c.ToHsv()
		expected := HSV{300, 100, 100}
		if hsv != expected {
			t.Errorf("expected %v, got %v", expected, hsv)
		}
	})

	t.Run("hsla to cmyk", func(t *testing.T) {
		c := HSLA{HSL{120, 100, 25}, 1.0} // dark green
		cmyk := c.ToCmyk()
		expected := CMYK{100, 0, 100, 50}
		if cmyk != expected {
			t.Errorf("expected %v, got %v", expected, cmyk)
		}
	})
}

func TestHsv(t *testing.T) {
	t.Run("string to hsv", func(t *testing.T) {
		hsv, _ := StrToHsv("hsv(0,100,100)")
		expected := HSV{0, 100, 100}
		if *hsv != expected {
			t.Errorf("expected %v, got %v", expected, *hsv)
		}
	})

	t.Run("hsv to string", func(t *testing.T) {
		c := HSV{240, 100, 100}
		str := c.String()
		expected := "hsv(240,100,100)"
		if str != expected {
			t.Errorf("expected %s, got %s", expected, str)
		}
	})

	t.Run("hsv to rgb", func(t *testing.T) {
		c := HSV{0, 0, 100} // white
		rgb := c.ToRgb()
		expected := RGB{255, 255, 255}
		if rgb != expected {
			t.Errorf("expected %v, got %v", expected, rgb)
		}
		
		// Test red
		redHsv := HSV{0, 100, 100}
		redRgb := redHsv.ToRgb()
		if redRgb != (RGB{255, 0, 0}) {
			t.Errorf("red conversion failed: %v", redRgb)
		}
	})

	t.Run("hsv to rgba", func(t *testing.T) {
		c := HSV{120, 100, 100} // green
		rgba := c.ToRgba()
		expected := RGBA{RGB{0, 255, 0}, 1.0}
		if rgba.R != expected.R || rgba.G != expected.G || 
		   rgba.B != expected.B || rgba.A != expected.A {
			t.Errorf("expected %v, got %v", expected, rgba)
		}
	})

	t.Run("hsv to hex", func(t *testing.T) {
		c := HSV{300, 100, 100} // magenta
		hex := c.ToHex()
		expected := "#ff00ff"
		if hex != expected {
			t.Errorf("expected %s, got %s", expected, hex)
		}
		
		// Test blue
		blueHsv := HSV{240, 100, 100}
		if blueHsv.ToHex() != "#0000ff" {
			t.Error("blue hex conversion failed")
		}
	})

	t.Run("hsv to hsl", func(t *testing.T) {
		c := HSV{60, 100, 100} // yellow
		hsl := c.ToHsl()
		expected := HSL{60, 100, 50}
		if hsl != expected {
			t.Errorf("expected %v, got %v", expected, hsl)
		}
		
		// Test black
		blackHsv := HSV{0, 0, 0}
		blackHsl := blackHsv.ToHsl()
		if blackHsl.L != 0 {
			t.Errorf("black conversion failed: %v", blackHsl)
		}
	})

	t.Run("hsv to hsla", func(t *testing.T) {
		c := HSV{0, 100, 50} // dark red
		hsla := c.ToHsla()
		expected := HSLA{HSL{0, 99, 24}, 1.0}
		if hsla.H != expected.H || hsla.S != expected.S || 
		   hsla.L != expected.L || hsla.A != expected.A {
			t.Errorf("expected %v, got %v", expected, hsla)
		}
	})

	t.Run("hsv to cmyk", func(t *testing.T) {
		c := HSV{0, 0, 0} // black
		cmyk := c.ToCmyk()
		expected := CMYK{0, 0, 0, 100}
		if cmyk != expected {
			t.Errorf("expected %v, got %v", expected, cmyk)
		}
		
		// Test cyan
		cyanHsv := HSV{180, 100, 100}
		cyanCmyk := cyanHsv.ToCmyk()
		if cyanCmyk.C != 100 || cyanCmyk.M != 0 || 
		   cyanCmyk.Y != 0 || cyanCmyk.K != 0 {
			t.Errorf("cyan conversion failed: %v", cyanCmyk)
		}
	})
}