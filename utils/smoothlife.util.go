package utils

import "math"

func SigmaSL(x, a, alpha float64) float64 {
	tempVal := -(x - a) * 4.0 / alpha
	return 1.0 / (1.0 + math.Exp(tempVal))
}

func SigmaNSL(x, a, b, alphaN float64) float64 {
	res := SigmaSL(x, a, alphaN) * (1 - SigmaSL(x, b, alphaN))

	return res
}

func SigmaMSL(x, y, m, alphaM float64) float64 {
	val1 := 1 - SigmaSL(m, 0.5, alphaM)
	val2 := SigmaSL(m, 0.5, alphaM)

	return x*val1 + y*val2
}

func SBigSL(b1, b2, d1, d2, m, n, alphaM, alphaN float64) float64 {
	val1 := SigmaMSL(b1, d1, m, alphaM)
	val2 := SigmaMSL(b2, d2, m, alphaM)

	return SigmaNSL(n, val1, val2, alphaN)
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	}

	return x
}
