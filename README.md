# Interpolating Polynomials

"Mathematics for Computer Programmers" chapter 2 Polynomials

# Example

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
