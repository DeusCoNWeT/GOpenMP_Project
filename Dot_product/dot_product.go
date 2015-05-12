package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum float64
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	a := make([]float64, n)
	b := make([]float64, n)
	var i int

	for i = 0; i < n; i++ {
		a[i] = float64(i) * 0.5
		b[i] = float64(i) * 2.0
	}
	sum = 0

	//pragma gomp parallel for private(i) reduction(+:sum)
	for i = 0; i < n; i++ {
		sum += a[i] * b[i]
	}
	fmt.Println("a*b =", sum)
}
