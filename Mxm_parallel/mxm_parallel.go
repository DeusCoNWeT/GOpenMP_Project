package main

import (
	"math"
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
	var _barrier_0_bool = make(chan bool)
	for _i := 0; _i < _numCPUs; _i++ {
		go func(_routine_num int, _angle float64) {
			var (
				angle float64
			)
			for i := _routine_num + 0; i < (n+0)/1; i += _numCPUs {
				var ()
				for j := 0; j < n; j++ {
					angle = 2.0 * pi * float64(i) * float64(j) / float64(n)
					a[i][j] = s * (math.Sin(angle) + math.Cos(angle))
				}
			}
			for i := _routine_num + 0; i < (n+0)/1; i += _numCPUs {
				var ()
				for j := 0; j < n; j++ {
					b[i][j] = a[i][j]
				}
			}
			for i := _routine_num + 0; i < (n+0)/1; i += _numCPUs {
				var ()
				for j := 0; j < n; j++ {
					c[i][j] = 0.0
					for k := 0; k < n; k++ {
						c[i][j] = c[i][j] + a[i][k]*b[k][j]
					}
				}
			}
			_barrier_0_bool <- true
		}(_i, angle)
	}
	for _i := 0; _i < _numCPUs; _i++ {
		<-_barrier_0_bool
	}

}
