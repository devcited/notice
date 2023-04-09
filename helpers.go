package notice

import (
	"math"
	"strconv"
)

func hexToRGB(hex string) (int, int, int) {
	if len(hex) == 4 {
		return hexToRGB(hex[0:1] + hex[1:2] + hex[1:2] + hex[2:3] + hex[2:3] + hex[3:4] + hex[3:4])
	}

	r, _ := strconv.ParseInt(hex[1:3], 16, 64)
	g, _ := strconv.ParseInt(hex[3:5], 16, 64)
	b, _ := strconv.ParseInt(hex[5:7], 16, 64)
	return int(r), int(g), int(b)
}

func hueToRGB(hue float64) (int, int, int) {
	hue = math.Mod(hue, 360)
	if hue < 0 {
		hue += 360
	}

	var r, g, b float64 = 0, 0, 0

	if hue < 60 {
		r, g, b = 1, hue/60, 0
	} else if hue < 120 {
		r, g, b = (120-hue)/60, 1, 0
	} else if hue < 180 {
		r, g, b = 0, 1, (hue-120)/60
	} else if hue < 240 {
		r, g, b = 0, (240-hue)/60, 1
	} else if hue < 300 {
		r, g, b = (hue-240)/60, 0, 1
	} else {
		r, g, b = 1, 0, (360-hue)/60
	}

	return int(r * 255), int(g * 255), int(b * 255)
}

func hslToRGB(hue, saturation, lightness int) (int, int, int) {
	h := float64(hue) / 360.0
	s := float64(saturation) / 100.0
	l := float64(lightness) / 100.0

	var r, g, b float64

	if s == 0 {
		r, g, b = l, l, l
	} else {
		q := l * (1 + s)
		if l >= 0.5 {
			q = l + s - (l * s)
		}
		p := 2*l - q

		r = hueToRGBHelper(p, q, h+1.0/3.0)
		g = hueToRGBHelper(p, q, h)
		b = hueToRGBHelper(p, q, h-1.0/3.0)
	}

	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func hueToRGBHelper(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}

	if t > 1 {
		t -= 1
	}

	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}

	if t < 1.0/2.0 {
		return q
	}

	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}

	return p
}
