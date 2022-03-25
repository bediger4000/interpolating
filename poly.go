package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, y []float64

	for _, str := range os.Args[1:] {
		fields := strings.Split(str, ",")
		xi, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			log.Printf("parsing %q: %v\n", fields[0], err)
			continue
		}
		yi, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Printf("parsing %q: %v\n", fields[1], err)
			continue
		}
		x = append(x, xi)
		y = append(y, yi)
	}

	for i := range x {
		ycalc := poly(x, y, x[i])
		fmt.Printf("at %f, calculated %f, wanted %f\n", x[i], ycalc, y[i])
	}
}

func poly(x, y []float64, z float64) float64 {
	sum := 0.0
	for i, yi := range y {
		product := 1.0
		for j, xj := range x {
			if i == j {
				continue
			}
			product *= (z - xj) / (x[i] - xj)
		}
		sum += yi * product
	}
	return sum
}
