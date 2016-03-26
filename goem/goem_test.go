package goem

import (
  "math"
  "testing"

  . "github.com/r7kamura/gospel"
)

func TestDescribe(t *testing.T) {
  Describe(t, "We initializes EM", func(){
	  matrix := [][]float64{{0, 0}, {0, 1}, {1, 0}, {3, 2}, {3, 3}, {2, 3}}
	  var em = NewEM(0.1, 3, matrix, 1)
		Context("and we confirm", func(){
		  It("should be the correct sigma", func(){
			mat := []float64{0.1, 0.1, 0.1}
			Expect(em.sigma).To(Equal, mat)
			})
		  It("should be the correct mu", func(){
		  	mu := [][]float64{{0, 0}, {1.0/3, 1.0/3}, {2.0/3, 2.0/3}}
			Expect(em.mu).To(Equal, mu)
			})
		})
		Context("and we calculate", func(){
		  em.EmIter(40, 0, true, "pic/")
		  It("should be the correct mu", func(){
			  for i := range em.mu {
				  for j := range em.mu[0] {
					  em.mu[i][j] = math.Floor(em.mu[i][j]*100)/100
					}
				}
			mu2 := [][]float64{{0, 0}, {0.49, 0.49}, {2.12, 2.12}}
			Expect(em.mu).To(Equal, mu2)
			})
		})
	})
}
