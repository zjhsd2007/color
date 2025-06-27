# Color Conversion Library

Go language-based color model conversion library that supports mutual conversion between multiple color formats, compliant with W3C color specifications.

## Core Features

- Supports RGB/RGBA, HEX, HSL/HSLA, HSV, CMYK color models
- Provides unified interface `Color` for color operations (see `interface.go`)
- Precision assurance: Color values use uint8/uint16, opacity uses float32
- Error handling: All parsing functions return `(result, error)`

## Installation

```bash
go get github.com/zjhsd2007/color
```

## Usage Example

```go
package main

import (
	"github.com/zjhsd2007/color"
)

func main() {
	hex, _ := color.StrToHex("#ff0000")
	rgb, _ := hex.ToRGB()
	rgbStr := rgb.String() // "rgb(255,0,0)"
}
```

## Design Characteristics

- Performance optimization: Avoid redundant calculations, integer operations take precedence over floating-point operations
- Extensibility: Modular design allows adding new color models by implementing base interface
- Test coverage: Comprehensive unit tests verify conversion accuracy (see `color_test.go`)
