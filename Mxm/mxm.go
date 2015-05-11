package main

import (
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	a := make([][]float64, n)
	b := make([][]float64, n)
	c := make([][]float64, n)
	for i := 0; i < len(a); i++ {
		a[i] = make([]float64, n)
		b[i] = make([]float64, n)
		c[i] = make([]float64, n)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			a[i][j] = 1
			b[i][j] = b[i][j] + a[i][j]
			c[i][j] = c[i][j] + b[i][j]
		}
	}

	var angle float64

	var pi float64 = 3.141592653589793
	var s float64

	s = 1.0 / math.Sqrt(float64(n))

	//pragma gomp parallel shared(a, b, c, n, pi, s) private(angle)
	{
		//pragma gomp for
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				angle = 2.0 * pi * float64(i) * float64(j) / float64(n)
				a[i][j] = s * (math.Sin(angle) + math.Cos(angle))
			}
		}
		//pragma gomp for
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				b[i][j] = a[i][j]
			}
		}
		//pragma gomp for
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				c[i][j] = 0.0
				for k := 0; k < n; k++ {
					c[i][j] = c[i][j] + a[i][k]*b[k][j]
				}
			}
		}
	}

}
