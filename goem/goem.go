package goem

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math"
)

//EM is EM-algorithm class.
type EM struct {
	mu    [][]float64
	sigma []float64
	d     int
	k     int
	w     [][]float64
	data  [][]float64 //x
	pi    []float64
}

//NewEM is initializer of struct EM.
func NewEM(sig float64, cluster int, data [][]float64) *EM {
	w := make([][]float64, len(data))
	for i := 0; i < len(data); i++ {
		w[i] = make([]float64, cluster)
	}
	pi := make([]float64, cluster) //stub.
	for i := range pi {
		pi[i] = 1.0 / float64(cluster)
	}
	mu := make([][]float64, cluster)
	for i := 0; i < cluster; i++ {
		mu[i] = make([]float64, len(data[0]))
	}
	sigma := make([]float64, cluster)
	for i := range sigma {
		sigma[i] = sig
	}
	EM := &EM{mu: mu, sigma: sigma, d: len(data[0]), k: cluster, data: data, pi: pi, w: w}
	EM.muInitAsBiasedMean(1.0)
	return EM
}

func (em EM) muInitAsBiasedMean(biase float64) {
	//em.mu = make([][]float64, em.k)
	mu := make([]float64, em.d)
	for i := 0; i < em.k; i++ {
		for d := 0; d < em.d; d++ {
			mu[d] += em.data[i][d]
		}
	}
	for i := 0; i < em.k; i++ {
		for d := 0; d < em.d; d++ {
			em.mu[i][d] = mu[d] / float64(em.k) * (1.0 + float64(i-em.k/2.0)*biase)
		}
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

	var jisuu = 0.5 * float64(em.d)
	return math.Exp(resultMat.At(0, 0)/(-2.0)/(em.sigma[j]*em.sigma[j])) / math.Pow(2*math.Pi*em.sigma[j]*em.sigma[j], jisuu)
}

func (em EM) e() {
	for n := 0; n < len(em.data); n++ {
		for k := 0; k < em.k; k++ {
			tmp := 0.0
			for j := 0; j < em.k; j++ {
				tmp += em.pi[j] * em.norm(em.data[n], j)
			}
			em.w[n][k] = em.pi[k] * em.norm(em.data[n], k) / tmp
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
			for d, v := range em.data[n] {
				biasedSigma += em.w[n][k] * math.Pow(v-em.mu[k][d], 2)
			}
			//biasedSigma += em.w[n][k] * arraySubInnerProduct(em.data[n], em.mu[k])
		}
		em.sigma[k] = math.Sqrt(biasedSigma / sumW[k] / float64(em.d))
		em.pi[k] = sumW[k] / float64(len(em.data))
	}
}

func (em EM) show() {
	for i := range em.pi {
		fmt.Println("pi", i, ": ", em.pi[i])
		fmt.Println("mu", i, ": ", em.mu[i])
		fmt.Println("sigma", i, ": ", em.sigma[i])
	}
	fmt.Println("loglikelyhood: ", em.likelyhood())
}

func arraySubInnerProduct(a []float64, b []float64) (result float64) {
	for i := range a {
		result += (a[i] - b[i]) * (a[i] - b[i])
	}
	return
}

func (em EM) EmIter(times int, loglikelyhood float64) {
	like := em.likelyhood()
	var i int
	for i = 0; i < times; i++ {
		em.show()
		em.e()
		em.m()
		newlike := em.likelyhood()
		if math.IsNaN(em.likelyhood()) { //|| math.Abs(newlike-like) < loglikelyhood {
			break
		}
		like = newlike
	}
	fmt.Println("iter: ", i, like)

}

func (em EM) likelyhood() (result float64) {
	for _, d := range em.data {
		temp := 0.0
		for k, v := range em.pi {
			temp += v * em.norm(d, k)
		}
		result += math.Log(temp)
	}
	return
}

/*
func main() {
	//mu := [][]float64{{0, 0}, {0, 0}, {0, 0}}
	data := [][]float64{{0.5, 0.2}, {0.4, 0.2}, {0.4, 0.3}, {0.3, 0.3}}
	a := NewEM(1.0, 3, data)
	a.emIter(5, 0.01)
}
*/
