package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/6br/goem/goem"
)

func ioinput() (data [][]float64) {
	for {
		var x float64
		var y float64
		i, err := fmt.Scan(&x, &y)
		if i == 2 {
			temp := []float64{x, y}
			data = append(data, temp)
		} else if i == 0 && err == io.EOF {
			break
		}
	}
	return
}

func main() {
	var (
		verbose      bool
		sigma        float64
		iter         int
		optimizeiter int
		clustermax   int
		partition    int
		loglikely    float64
		meanshift    float64
		goroutine    bool
	)
	/* register flag name and shorthand name */
	flag.BoolVar(&verbose, "verbose", true, "verbose option(show midstream and plot graph)")
	flag.BoolVar(&verbose, "v", true, "verbose option(show midstream and plot graph)")
	flag.BoolVar(&goroutine, "g", true, "use goroutine for optimize or not")
	flag.BoolVar(&goroutine, "goroutne", true, "use goroutine for optimize or not")
	flag.Float64Var(&sigma, "sigma", 1.0, "initial sample variance")
	flag.Float64Var(&sigma, "s", 1.0, "initial sample variance")
	flag.Float64Var(&meanshift, "meanshift", 1.0, "how many shift from mean of data when mu set")
	flag.Float64Var(&meanshift, "m", 1.0, "how many shift from mean of data when mu set")
	flag.Float64Var(&loglikely, "loglikelyhood", 0.01, "permit the highest loglikelyhood for EM-algorithm")
	flag.Float64Var(&loglikely, "l", 0.01, "permit the highest loglikelyhood for EM-algorithm")
	flag.IntVar(&iter, "i", 40, "iteration for precise distermination of cluster center")
	flag.IntVar(&iter, "iter", 40, "iteration for precise distermination of cluster center")
	flag.IntVar(&optimizeiter, "o", 20, "iteration for about search for cluster's number")
	flag.IntVar(&optimizeiter, "optimizeiter", 20, "iteration for about search for cluster's number")
	flag.IntVar(&clustermax, "c", 7, "max cluster size for optimizing parameter")
	flag.IntVar(&clustermax, "clustermax", 7, "max cluster size for optimizing parameter")
	flag.IntVar(&partition, "p", 5, "separate data for calculate cross-entropy")
	flag.IntVar(&partition, "partition", 5, "separate data for calculate cross-entropy")
	flag.Parse()

	//fmt.Println(data)
	//data := [][]float64{{0.5, 0.2}, {0.4, 0.2}, {0.4, 0.3}, {0.3, 0.3}}
	//a := goem.NewEM(1, 3, data)
	//a.EmIter(iter, loglikely)
	//b := goem.NewEM(sigma, 3, ioinput(), 1.0)
	//b.CrossEntropy(7, 5, 20, 1.0)
	//b.EmIter(iter, loglikely, verbose)
	//b.Show()
	em := goem.NewOptimizedEM(sigma, clustermax, partition, optimizeiter, ioinput(), meanshift, goroutine)
	em.EmIter(iter, loglikely, verbose)
	//fmt.Println(em.CrossEntropy(10, 5, 20)) // from 2 to 9, partition by 5.
	//crossem := goem.NewEM(1, em.CrossEntropy(10, 5, 20), em.data)
}
