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
		directory    string
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
	flag.StringVar(&directory, "d", "", "path to the directory to save plotted images(ex: pic/)")
	flag.StringVar(&directory, "directory", "", "path to the directory to save plotted images(ex: pic/)")
	flag.Parse()

	em := goem.NewOptimizedEM(sigma, clustermax, partition, optimizeiter, ioinput(), meanshift, goroutine)
	em.EmIter(iter, loglikely, verbose, directory)
	fmt.Println("Result: ")
	em.Show()
}
