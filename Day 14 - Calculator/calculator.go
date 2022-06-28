package main

var Operations = map[int]func(float64, float64) float64{
	1: func(i1, i2 float64) float64 {
		return i1 + i2
	},
	2: func(i1, i2 float64) float64 {
		return i1 - i2
	},
	3: func(i1, i2 float64) float64 {
		return i1 / i2
	},
	4: func(i1, i2 float64) float64 {
		return i1 * i2
	},
}
