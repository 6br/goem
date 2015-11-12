package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	elemX := []float64{1, 2, 3, 4}
	elemY := []float64{5, 6, 7, 8}
	elemA := make([]float64, 4)

	x := mat64.NewDense(2, 2, elemX)
	y := mat64.NewDense(2, 2, elemY)
	a := mat64.NewDense(2, 2, elemA)

	a.Mul(x, y)
	fmt.Println(x)
}
