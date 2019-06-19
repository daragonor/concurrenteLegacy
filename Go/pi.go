package main

import (
	"fmt"
	"time"
)

const N = 100000000

func calcPi() float64 {
	var x, step, sum float64
	step = 1.0 / N
	sum = 0
	for i := 0; i < N; i++ {
		x = (float64(i) + 0.5) * step
		sum += 4.0 / (1 + x*x)
	}
	return step * sum
}

func main() {
	ini := time.Now()
	fmt.Println(calcPi())
	fmt.Printf("Tiempo: %s\n", time.Since(ini))
}
