package goem

import (
	"fmt"
	"math"
)

//BinaryParams is needed to use goroutine.
type BinaryParams struct {
	params int
	score  float64
}

//CrossEntropy calculates the best cluster number.
func (em EM) CrossEntropy(end int, partition int, iter int, mean float64) (bestCluster int) {
	bestEntropy := math.MaxFloat64
	for i := 2; i < end; i++ {
		entropy := 0.0
		for part := 0; part < partition; part++ {
			trainData, testData := em.crossEM(part*len(em.data)/partition, (part+1)*len(em.data)/partition)
			crossem := NewEM(1.0, i, trainData, mean)
			crossem.EmIter(iter, 0.01, false, "") //loosen constraint.
			entropy += crossem.entropy(testData)
		}
		entropy /= float64(i)
		fmt.Println("Clusters:", i, " Entropy: ", entropy)
		if entropy < bestEntropy {
			bestEntropy = entropy
			bestCluster = i
		}
	}
	return
}

func (em EM) parallelCrossEntropy(end int, partition int, iter int, mean float64) int {
	//Generate 0-1 array and each throw go-routine
	ch := make(chan BinaryParams)
	for i := end; i >= 2; i-- {
		go em.eachCrossEntropy(end, partition, iter, mean, i, ch)
	}
	//Gather results and return most optimal solution
	var scoremax BinaryParams
	scoremax.score = math.MaxFloat64
	for i := 2; i < end; i++ {
		params := <-ch
		if params.score < scoremax.score {
			scoremax = params
		}
	}
	return scoremax.params
}

func (em EM) eachCrossEntropy(end int, partition int, iter int, mean float64, i int, ch chan BinaryParams) {
	entropy := 0.0
	for part := 0; part < partition; part++ {
		trainData, testData := em.crossEM(part*len(em.data)/partition, (part+1)*len(em.data)/partition)
		crossem := NewEM(1.0, i, trainData, mean)
		crossem.EmIter(iter, 0.01, false,"") //loosen constraint.
		entropy += crossem.entropy(testData)
	}
	entropy /= float64(i)
	fmt.Println("Clusters:", i, " Entropy: ", entropy)
	ch <- BinaryParams{i, entropy}
}

//NewOptimizedEM generates EM object with optimized param by cross-entropy
func NewOptimizedEM(sig float64, end int, partition int, iter int, data [][]float64, mean float64, goroutine bool) *EM {
	em := NewEM(sig, end, data, mean)
	if end > len(data)/partition * (partition-1) {
	  end = len(data)/partition * (partition-1)
	}
	var newem *EM
	if goroutine {
		newem = NewEM(sig, em.parallelCrossEntropy(end, partition, iter, mean), data[:], mean)
	} else {
		newem = NewEM(sig, em.CrossEntropy(end, partition, iter, mean), data[:], mean)
	}
	return newem
}

func (em EM) crossEM(start int, end int) (trainData [][]float64, testData [][]float64) {
	testData = em.data[start:end]
	trainData2 := em.data[:start]
	trainData3 := em.data[end:]
	trainData = append(trainData, trainData2...)
	trainData = append(trainData, trainData3...)
	return
}

func (em EM) entropy(testData [][]float64) float64 {
	entropy := 0.0
	for _, v := range testData {
		temp := 0.0
		for k := 0; k < em.k; k++ {
			temp += em.pi[k] * em.norm(v, k)
		}
		entropy += math.Log(temp)
	}
	return entropy / -float64(len(testData))
}
