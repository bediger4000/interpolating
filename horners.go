package main

/*
 * Benchmark polynomial evaluation via Horner's method
 * (https://en.wikipedia.org/wiki/Horner%27s_method)
 r compared to a more literal method.
*/

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const (
	FROM    = -50.
	TO      = 50.
	DELTA   = .01
	EPSILON = 0.00001
)

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	coefficients := make([]float64, 0, len(os.Args))
	for _, str := range os.Args[1:] {
		c, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Printf("parsing %q: %v\n", str, err)
			continue
		}
		coefficients = append(coefficients, c)
	}
	fmt.Printf("# coefficients %v\n", coefficients)
	evaluationCount := 0
	for x := FROM; x <= TO; x += DELTA {
		a := hornersMethod(x, coefficients)
		b := literalMethod(x, coefficients)
		evaluationCount++
		if math.Abs(a-b) > 0.00001 {
			fmt.Printf("large diff at %f: %f - %f == %f\n",
				x, a, b, a-b)
		}
	}
	fmt.Printf("Evaluate polynomials %d times\n", evaluationCount)

	before := time.Now()
	for x := FROM; x <= TO; x += DELTA {
		hornersMethod(x, coefficients)
	}
	etHorners := time.Since(before)
	before = time.Now()
	for x := FROM; x <= TO; x += DELTA {
		literalMethod(x, coefficients)
	}
	etStandard := time.Since(before)

	fmt.Printf("Horner's method: %v\n", etHorners)
	fmt.Printf("Standard method: %v\n", etStandard)
}

// literalMethod evaluates each term from the highest degree
// to the lowest as they arrive.
func literalMethod(x float64, coefficients []float64) float64 {
	ln := len(coefficients)

	sum := 0.0

	for i := 0; i < ln; i++ {
		factor := coefficients[i]
		for power := ln - i; power > 1; power-- {
			factor *= x
		}
		sum += factor
	}

	return sum
}

// hornersMethod evaluates a polynomial at some value x where the x gets
// repeatedly factored out of parts of the polynomial:
// ax^3 + bx^x + cx +d  = [(ax + b)*x + c]x + d
func hornersMethod(x float64, coefficients []float64) float64 {
	ln := len(coefficients)
	sum := coefficients[0]
	for i := 1; i < ln; i++ {
		sum = sum*x + coefficients[i]
	}
	return sum
}

func usage() {
	fmt.Fprint(os.Stderr,
		`Benchmark polynomial evaluation, comparing Horner's method to a more literal
arithmetic evaluation of a normal form polynomial.
`)
	fmt.Fprint(os.Stderr, "usage: ./horners aN aM ... a0, where N is the degree of the polynomial\n")
}
