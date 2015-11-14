package main

import (
	"github.com/6br/goem/goem"
)

func main() {
	data := [][]float64{{0.5, 0.2}, {0.4, 0.2}, {0.4, 0.3}, {0.3, 0.3}}
	a := goem.NewEM(1.0, 3, data)
	a.EmIter(5, 0.01)
}
