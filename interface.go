package color

type String interface {
	String() string
}

type ToRgb interface {
	ToRgb() RGB
}

type ToRgba interface {
	ToRgba() RGBA
}

type ToHex interface {
	ToHex() string
}

type ToHsl interface {
	ToHsl() HSL
}

type ToHsla interface {
	ToHsla() HSLA
}

type ToHsv interface {
	ToHsv() HSV
}

type ToCmyk interface {
	ToCmyk() CMYK
}