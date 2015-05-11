package main

import (
	"flag"
)

import "runtime"

var (
	n_factor	= flag.Int("n_factor", 2, "Input a number")
	n_hi		= flag.Int("n_hi", 131072, "Input a number")
	n_lo		= flag.Int("n_lo", 1, "Input a number")
)
var _numCPUs = runtime.NumCPU()

func _init_numCPUs() {
	runtime.GOMAXPROCS(_numCPUs)
}
func prime_number(n int) int {
	var i int
	var j int
	var prime int
	var total int = 0
	var _barrier_0_int = make(chan int)
	for _i := 0; _i < _numCPUs; _i++ {
		go func(_routine_num int, _i int, _j int, _prime int) {
			var (
				i	int
				j	int
				prime	int
			)
			var total int
			for i = _routine_num + 2; i < (n+1)/1; i += _numCPUs {
				prime = 1
				for j = 2; j < i; j++ {
					if i%j == 0 {
						prime = 0
						break
					}
				}
				total = total + prime
			}
			_barrier_0_int <- total
		}(_i, i, j, prime)
	}
	for _i := 0; _i < _numCPUs; _i++ {
		total += <-_barrier_0_int
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
	_init_numCPUs()
	flag.Parse()
	prime_number_sweep(*n_lo, *n_hi, *n_factor)
}
