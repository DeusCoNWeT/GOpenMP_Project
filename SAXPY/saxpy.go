package main

import (
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
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	a := 5.0
	x := make([]float64, n)
	y := make([]float64, n)
	var _barrier_0_bool = make(chan bool)
	for _i := 0; _i < _numCPUs; _i++ {
		go func(_routine_num int) {
			var ()
			for i := _routine_num * (n / _numCPUs); i < (_routine_num+1)*(n/_numCPUs); i++ {

				y[i] = a*x[i] + y[i]
			}
			_barrier_0_bool <- true
		}(_i)
	}
	for _i := 0; _i < _numCPUs; _i++ {
		<-_barrier_0_bool
	}

}
