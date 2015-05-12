package main

import (
	"fmt"
	"os"
	"strconv"
)

import "runtime"

var _numCPUs = runtime.NumCPU()

func _init_numCPUs() {
	runtime.GOMAXPROCS(_numCPUs)
}
func main() {
	_init_numCPUs()
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
	var _barrier_0_float64 = make(chan float64)
	for _i := 0; _i < _numCPUs; _i++ {
		go func(_routine_num int, _i int) {
			var (
				i int
			)
			var sum float64
			for i = _routine_num + 0; i < (n+0)/1; i += _numCPUs {
				sum += a[i] * b[i]
			}
			_barrier_0_float64 <- sum
		}(_i, i)
	}
	for _i := 0; _i < _numCPUs; _i++ {
		sum += <-_barrier_0_float64
	}

	fmt.Println("a*b =", sum)
}
