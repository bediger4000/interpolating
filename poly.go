package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	min := flag.Float64("min", 0.0, "minimum domain value")
	max := flag.Float64("max", 10.0, "maximum domain value")
	step := flag.Float64("step", 0.01, "domain delta value")
	flag.Parse()

	var x, y []float64

	n := flag.NArg()

	for i := 0; i < n; i++ {
		str := flag.Arg(i)
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

	for i, d := range x {
		r := poly(x, y, d)
		delta := y[i] - r
		if delta > 0.001 {
			fmt.Fprintf(os.Stderr, "%f wanted %f got %f (%f)\n", d, r, y[i], delta)
		}
	}

	for r := *min; r <= *max; r += *step {
		ycalc := poly(x, y, r)
		fmt.Printf("%f\t%f\n", r, ycalc)
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
