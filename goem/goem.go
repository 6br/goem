package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math"
)

//EM is EM-algorithm class.
type EM struct {
	mu    [][]float64
	sigma float64
	d     int
	k     int
	w     [][]float64
	data  [][]float64 //x
	pi    []float64
}

//NewEM is initializer of struct EM.
func NewEM(mu [][]float64, sigma float64, cluster int, data [][]float64) *EM {
	w := make([][]float64, len(data))
	for i := 0; i < len(data); i++ {
		w[i] = make([]float64, cluster)
	}
	//pi := []float64{0.3, 0.3, 0.3}
	pi := make([]float64, cluster) //stub.
	for i := range pi {
		pi[i] = 1.0 / float64(cluster)
	}
	EM := &EM{mu: mu, sigma: sigma, d: len(data), k: cluster, data: data, pi: pi, w: w}
	return EM
}

func (em EM) muInitAsBiasedMean() {
	mu := make([][]float64, em.k)
	for i := 0; i < em.k; i++ {
		mu[i] = make([]float64, em.d)
	}
}

func (em EM) norm(x []float64, j int) float64 {
	xMat := mat64.NewDense(1, len(x), x)
	muMat := mat64.NewDense(1, len(em.mu[j]), em.mu[j])
	first := mat64.NewDense(1, len(em.mu[j]), nil)
	first.Sub(xMat, muMat)
	second := mat64.DenseCopyOf(first.T())
	resultMat := mat64.NewDense(1, 1, nil)
	resultMat.Mul(first, second)
	//fmt.Println(resultMat.At(0, 0))
	var jisuu = 0.5 * float64(em.d)
	return math.Exp(resultMat.At(0, 0)/(-2.0)/(em.sigma*em.sigma)) / math.Pow(2*math.Pi*em.sigma*em.sigma, jisuu)
}

func (em EM) e() {
	for n := 0; n < len(em.data); n++ {
		for k := 0; k < em.k; k++ {
			tmp := 0.0
			for j := 0; j < em.k; j++ {
				tmp += em.pi[j] * em.norm(em.data[n], j)
			}
			em.w[n][k] = em.pi[k] * em.norm(em.data[n], k)
		}
	}
}

func (em EM) m() {
	sumW := make([]float64, em.k)
	for k := 0; k < em.k; k++ {
		sumW[k] = 0.0
		for n := 0; n < len(em.data); n++ {
			sumW[k] += em.w[n][k]
		}
	}

	for k := 0; k < em.k; k++ {
		biasedMu := make([]float64, em.d)
		for n := 0; n < len(em.data); n++ {
			for d, v := range em.data[n] {
				biasedMu[d] += em.w[n][k] * v
			}
		}
		for f, v := range biasedMu {
			em.mu[k][f] = v / sumW[k]
		}

		biasedSigma := 0.0
		for n := 0; n < len(em.data); n++ {
			biasedSigma += em.w[n][k] * math.Pow(arraySubInnerProduct(em.data[n], em.mu[k]), 2)
		}
		em.sigma = biasedSigma / sumW[k]
		em.pi[k] = sumW[k] / float64(len(em.data))
	}
}

func arraySubInnerProduct(a []float64, b []float64) (result float64) {
	result = 0
	for i := range a {
		result += (a[i] - b[i]) * (a[i] - b[i])
	}
	return
}

func main() {
	mu := [][]float64{{0, 0}, {0, 0}, {0, 0}}
	x := []float64{1, 2}
	data := [][]float64{{0.5, 0.2}, {0.4, 0.2}, {0.4, 0.2}}
	a := NewEM(mu, 1.0, 3, data)
	fmt.Println(a.norm(x, 0))
	a.e()
	a.m()
}
