package main

import (
	"os"
	"strconv"
)

func main() {
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	a := 5.0
	x := make([]float64, n)
	y := make([]float64, n)

	//pragma gomp parallel for
	for i := 0; i < n; i++ {

		y[i] = a*x[i] + y[i]
	}

}
