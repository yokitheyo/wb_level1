package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	grp := make(map[int][]float64)

	for _, t := range temps {
		k := int(math.Floor(t/10) * 10)
		grp[k] = append(grp[k], t)
	}

	for k, v := range grp {
		fmt.Printf("%d: %v\n", k, v)
	}
}
