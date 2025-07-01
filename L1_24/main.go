package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Disntance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)
	distance := p1.Disntance(*p2)
	fmt.Printf("distance between p1 and p2: %.4f\n", distance)
}
