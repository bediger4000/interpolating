# Interpolating Polynomials

[Programmer's Introduction to Mathematics](https://pimbook.org/),
chapter 2, "Polynomials".

That chapter present polynomials as an example
mathematical object.
It proves that you can construct a unique polynomial
that passes through N points on the plane.

This repo exists to prove to myself that I understand the ideas.

## Example

```sh
$ ./poly -min -10 -max 10  0,0 1,1 -1,1 > data
$ ./poly -min -10 -max 10  5,0 6,1 4,1 > data2
```
Should output numerical values for y = x<sup>2</sup>
and y = x<sup>2</sup> + 5, two concave-upward parabolas.

Since 3 (N=3) points get interpolated,
the program will calculate a polynomial of degree 2 (N - 1)

[gnuplot](http://www.gnuplot.info/)
could plot those two sets of data with
a command like this:

```
gnuplot> set xrange [-10:20]
gnuplot> plot 'data' with lines, 'data2' with lines
```

## Bonus Horner's Method Benchmarking

[Code](horners.go)

One of the exercises for Chapter 1 is to "look up Horner's method"
and benchmark polynomial evaluation versus some other,
more obvious method.

My program evaluates a polynomial (coefficients given on the command line)

```sh
$ ./horners -.00001 0.0 -.0001 .053 12 12020.
```

would evalute a 5th-degree polynomial:

-0.00001x<sup>5</sup> - 0.0001x<sup>3</sup> +.053x<sup>2</sup> + 12x + 12020

The code contains constants for range of X, and delta-X
so as to execute each method of polynomial evaluate a large number of times.
