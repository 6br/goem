package goem

import (
	"fmt"
	"math"
)

//CrossEntropy culculates best bluster number.
func (em EM) CrossEntropy(end int, partition int, iter int, mean float64) (bestCluster int) {
	bestEntropy := 0.0
	for i := 2; i < end; i++ {
		entropy := 0.0
		for part := 0; part < partition; part++ {
			trainData, testData := em.crossEM(part*len(em.data)/partition, (part+1)*len(em.data)/partition)
			crossem := NewEM(1.0, i, trainData, mean)
			crossem.EmIter(iter, 0.01, false) //緩めの条件で回す。
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

//NewOptimizedEM generates EM object with optimized param by cross-entropy
func NewOptimizedEM(sig float64, end int, partition int, iter int, data [][]float64, mean float64) *EM {
	em := NewEM(sig, end, data, mean)
	//em.recluster(em.CrossEntropy(end, partition, iter, mean), sig)
	newem := NewEM(sig, em.CrossEntropy(end, partition, iter, mean), data[:], mean)
	return newem
}

//TODO Use goroutine
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
