package main

import (
	"flag"
)

var (
	n_factor = flag.Int("n_factor", 2, "Input a number")
	n_hi     = flag.Int("n_hi", 131072, "Input a number")
	n_lo     = flag.Int("n_lo", 1, "Input a number")
)

func prime_number(n int) int {
	var i int
	var j int
	var prime int
	var total int = 0

	//pragma gomp parallel for shared(n) private(i, j, prime) reduction (+:total)
	for i = 2; i <= n; i++ {
		prime = 1
		for j = 2; j < i; j++ {
			if i%j == 0 {
				prime = 0
				break
			}
		}
		total = total + prime
	}
	return total
}

func prime_number_sweep(n_lo int, n_hi int, n_factor int) {
	var n int

	n = n_lo

	for n <= n_hi {

		_ = prime_number(n)

		n = n * n_factor
	}

}

func main() {

	flag.Parse()

	prime_number_sweep(*n_lo, *n_hi, *n_factor)

}
