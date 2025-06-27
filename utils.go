package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calcRgbWithAlpha(v uint8, alpha float32) uint8 {
    return clampToUint8(float32(v) * alpha + 255 * (1 - alpha))
}

func rgbToHsl(r uint8, g uint8, b uint8) (uint32, uint32, uint32){
  return calcRgbToHsl(float32(r) / 255.0, float32(g)/ 255.0, float32(b)/ 255.0)
}

func hslToRgb(h uint32, s uint32, l uint32)(uint8, uint8, uint8) {
  sNorm := float64(s) / 100.0
  lNorm := float64(l) / 100.0
  c := (1.0 - math.Abs(2.0*lNorm-1.0)) * sNorm
  x := c * (1.0 - math.Abs(math.Mod(float64(h)/60.0, 2.0)-1.0))
  m := lNorm - c/2.0
  
  var r, g, b float64
  
  switch {
    case h < 60:
      r, g, b = c, x, 0.0
    case h >= 60 && h < 120:
      r, g, b = x, c, 0.0
    case h >= 120 && h < 180:
      r, g, b = 0.0, c, x
    case h >= 180 && h < 240:
      r, g, b = 0.0, x, c
    case h >= 240 && h < 300:
      r, g, b = x, 0.0, c
    case h >= 300 && h < 360:
      r, g, b = c, 0.0, x
    default:
      r, g, b = 0.0, 0.0, 0.0
  }
  
  r = (r + m) * 255.0
  g = (g + m) * 255.0
  b = (b + m) * 255.0
  
  return uint8(math.Round(r)), uint8(math.Round(g)), uint8(math.Round(b))

}

func rgbToHsv(r, g, b uint8) (uint32, uint32, uint32) {
  rf := float64(r) / 255.0
  gf := float64(g) / 255.0
  bf := float64(b) / 255.0

  cMax := math.Max(math.Max(rf, gf), bf)
  cMin := math.Min(math.Min(rf, gf), bf)
  delta := cMax - cMin

  var h float64

  switch {
  case delta == 0:
    h = 0
  case cMax == rf:
    h = 60 * math.Mod((gf-bf)/delta, 6)
  case cMax == gf:
    h = 60 * ((bf-rf)/delta + 2)
  case cMax == bf:
    h = 60 * ((rf-gf)/delta + 4)
  }

  if h < 0 {
    h += 360
  }

  s := 0.0
  if cMax != 0 {
    s = delta / cMax
  }

  v := cMax

  return uint32(math.Round(h)), 
         uint32(math.Round(s * 100)), 
         uint32(math.Round(v * 100))
}

func hsvToRgb(h uint32, s uint32, v uint32) (uint8, uint8, uint8) {
  sNorm := float64(s) / 100.0
  vNorm := float64(v) / 100.0
  
  c := vNorm * sNorm
  x := c * (1 - math.Abs(math.Mod(float64(h)/60.0, 2)-1))
  m := vNorm - c
  
  var r, g, b float64
  
  switch {
  case h < 60:
    r, g, b = c, x, 0.0
  case h >= 60 && h < 120:
    r, g, b = x, c, 0.0
  case h >= 120 && h < 180:
    r, g, b = 0.0, c, x
  case h >= 180 && h < 240:
    r, g, b = 0.0, x, c
  case h >= 240 && h < 300:
    r, g, b = x, 0.0, c
  case h >= 300 && h < 360:
    r, g, b = c, 0.0, x
  default:
    r, g, b = 0.0, 0.0, 0.0
  }
  
  r = (r + m) * 255.0
  g = (g + m) * 255.0
  b = (b + m) * 255.0
  
  return uint8(r), uint8(g), uint8(math.Round(b))
}


func rgbToCmyk(r , g , b uint8) (uint8, uint8, uint8, uint8) {
  rf := float64(r) / 255.0
  gf := float64(g) / 255.0
  bf := float64(b) / 255.0

  // 计算K（黑色）分量
  k := 1.0 - math.Max(math.Max(rf, gf), bf)

  var c, m, y float64
  if k == 1.0 {
    // 纯黑情况
    c, m, y = 0.0, 0.0, 0.0
  } else {
    // 计算CMY分量
    c = (1.0 - rf - k) / (1.0 - k)
    m = (1.0 - gf - k) / (1.0 - k)
    y = (1.0 - bf - k) / (1.0 - k)
  }

  // 转换为百分比并四舍五入
  return uint8(math.Round(c * 100)), uint8(math.Round(m * 100)), uint8(math.Round(y * 100)), uint8(math.Round(k * 100))
}

func cmykToRgb(c uint8, m uint8, y uint8, k uint8) (uint8, uint8, uint8) {
  // 将CMYK百分比转换为0-1范围的浮点数
  cf := float64(c) / 100.0
  mf := float64(m) / 100.0
  yf := float64(y) / 100.0
  kf := float64(k) / 100.0

  // 计算转换因子
  t := 1.0 - kf

  // 计算RGB分量并四舍五入
  r := math.Round((1.0 - cf) * t * 255.0)
  g := math.Round((1.0 - mf) * t * 255.0)
  b := math.Round((1.0 - yf) * t * 255.0)

  // 确保结果在0-255范围内
  r = math.Max(0, math.Min(255, r))
  g = math.Max(0, math.Min(255, g))
  b = math.Max(0, math.Min(255, b))

  return uint8(r), uint8(g), uint8(b)
}

func rgbaToHsla(r uint8, g uint8, b uint8, a float32) (uint32, uint32, uint32, float32){
  h,s,l := rgbToHsl(r,g,b)
  return h,s,l,a
}

// hexToRGBA 将HEX颜色代码转换为RGBA值
// 支持格式: #rgb, #rgba, #rrggbb, #rrggbbaa
// 返回: r, g, b, a (每个分量范围0-255)
func hexToRGBA(hex string) (uint8, uint8, uint8, uint8, error) {
	// 去除可能的前后空白和#号
	hex = strings.TrimSpace(hex)
	if strings.HasPrefix(hex, "#") {
		hex = hex[1:]
	}

	var r, g, b, a uint8
	a = 255 // 默认完全不透明

	switch len(hex) {
	case 3: // #rgb 格式
		// 扩展为6字符格式 #rrggbb
		hex = string([]byte{
			hex[0], hex[0],
			hex[1], hex[1],
			hex[2], hex[2],
		})
		fallthrough
	case 6: // #rrggbb 格式
		// 解析RGB分量
		var err error
		r, err = parseHexComponent(hex[0:2])
		if err != nil {
			return 0, 0, 0, 0, err
		}
		g, err = parseHexComponent(hex[2:4])
		if err != nil {
			return 0, 0, 0, 0, err
		}
		b, err = parseHexComponent(hex[4:6])
		if err != nil {
			return 0, 0, 0, 0, err
		}
	case 4: // #rgba 格式
		// 扩展为8字符格式 #rrggbbaa
		hex = string([]byte{
			hex[0], hex[0],
			hex[1], hex[1],
			hex[2], hex[2],
			hex[3], hex[3],
		})
		fallthrough
	case 8: // #rrggbbaa 格式
		// 解析RGBA分量
		var err error
		r, err = parseHexComponent(hex[0:2])
		if err != nil {
			return 0, 0, 0, 0, err
		}
		g, err = parseHexComponent(hex[2:4])
		if err != nil {
			return 0, 0, 0, 0, err
		}
		b, err = parseHexComponent(hex[4:6])
		if err != nil {
			return 0, 0, 0, 0, err
		}
		a, err = parseHexComponent(hex[6:8])
		if err != nil {
			return 0, 0, 0, 0, err
		}
	default:
		return 0, 0, 0, 0, fmt.Errorf("invalid hex color format: %s", hex)
	}

	return r, g, b, a, nil
}

// parseHexComponent 解析两位十六进制字符串为uint8
func parseHexComponent(comp string) (uint8, error) {
	val, err := strconv.ParseUint(comp, 16, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid hex component: %s", comp)
	}
	return uint8(val), nil
}

/***************************************************************************************/
func clampToUint8(v float32) uint8 {
  if v < 0 {
    return 0
  }
  if v > 255 {
    return 255
  }
  return uint8(v + 0.5)
}

func absFloat32(x float32) float32 {
  if x < 0 {
    return -x
  }
  return x
}

func calcRgbToHsl(r float32, g float32, b float32)(uint32, uint32, uint32){
  cMax := max(r, g, b)
  cMin := min(r, g, b)
  delta := cMax - cMin
  h := calcHue(delta, cMax, r, g, b)
  if h < 0 {
    h += 360
  }
  l := (cMax + cMin) / 2
  var s float32
  if delta == 0 {
    s = 0
  } else {
    s = delta / (1.0 - absFloat32(2.0*l - 1.0))
  }
  return uint32(h), uint32(s * 100), uint32(l * 100)
}

func calcHue(delta float32, cMax float32, r float32, g float32, b float32) float32 {
  if delta == 0 {
    return 0
  }
  if cMax == r {
    return 60.0 * (g - b) / delta
  }
  if cMax == g {
    return 60.0 * (b - r) / delta + 120.0
  }
  return 60.0 * (r - g) / delta + 240.0
}

func RemoveSpace(str string) string {
  return strings.ReplaceAll(str, " ", "")
}