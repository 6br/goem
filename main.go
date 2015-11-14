package main

import (
	"fmt"
	"github.com/6br/goem/goem"
	"io"
	"os"
)

//TODO(6br) refer to http://qiita.com/ikawaha/items/28186d965780fab5533d
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
	//a := goem.NewEM(1, 3, data)
	//a.EmIter(5, 0.01)
	b := goem.NewEM(1, 3, ioinput(), 0.1)
	//b.CrossEntropy(7, 5, 20, 1.0)
	b.EmIter(30, 0.01, true)
	//b.Show()
	//em := goem.NewOptimizedEM(1, 7, 5, 20, ioinput(), 0.1)
	//em.EmIter(40, 0.01, true)
	//fmt.Println(em.CrossEntropy(10, 5, 20)) // from 2 to 9, partition by 5.
	//crossem := goem.NewEM(1, em.CrossEntropy(10, 5, 20), em.data)
}
