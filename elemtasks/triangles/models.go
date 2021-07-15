package main

import "math"

type triangle struct {
	Name  string
	Side1 float64
	Side2 float64
	Side3 float64
	Area  float64
}

func (t *triangle) isTriangle() (b bool) {
	if t.Side1+t.Side2 > t.Side3 && t.Side1+t.Side3 > t.Side2 && t.Side2+t.Side3 > t.Side1 {
		b = true
	}
	return
}

func (t *triangle) setHeronArea() {
	p := (t.Side1 + t.Side2 + t.Side3) / 2
	t.Area = math.Sqrt(p * (p - t.Side1) * (p - t.Side2) * (p - t.Side3))
}
