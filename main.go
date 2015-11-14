package main

import (
	"fmt"
	"github.com/6br/goem/goem"
	"io"
	"os"
)

func ioinput() (data [][]float64) {
	//data := make([][]float64, 0)
	for {
		var x float64
		var y float64
		i, err := fmt.Scan(&x, &y)
		if i == 2 {
			temp := []float64{x, y}
			data = append(data, temp)
		} else if i == 0 && err == io.EOF {
			break
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	return
}

func main() {
	//data := [][]float64{{0.5, 0.2}, {0.4, 0.2}, {0.4, 0.3}, {0.3, 0.3}}
	data := [][]float64{{0.5, 0.2}, {0.4, 0.2}, {0.4, 0.3}, {0.3, 0.3}}
	a := goem.NewEM(10, 3, data)
	a.EmIter(5, 0.01)
	b := goem.NewEM(10, 2, ioinput())
	b.EmIter(20, 0.00001)
	b.Plot()
}
