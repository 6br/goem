
package goem

import (
  "testing"
  . "github.com/r7kamura/gospel"
)

func TestDescribe2(t *testing.T) {
  Describe(t, "We initializes EM with goroutine", func(){
	  matrix := [][]float64{{0, 0}, {0, 1}, {1, 0}, {3, 2}, {3, 3}, {2, 3}}
	  var em = NewOptimizedEM(0.1, 4, 2, 20, matrix, 1, true)
		Context("and we confirm", func(){
		  It("should be the correct sigma", func(){
			  mat := []float64{0.1, 0.1}
			  Expect(em.sigma).To(Equal, mat)
			})
		})
	})
  Describe(t, "We initializes EM without goroutine", func(){
	  matrix := [][]float64{{0, 0}, {0, 1}, {1, 0}, {3, 2}, {3, 3}, {2, 3}}
	  var em = NewOptimizedEM(0.1, 4, 2, 20, matrix, 1, false)
		Context("and we confirm", func(){
		  It("should be the correct sigma", func(){
			  mat := []float64{0.1, 0.1}
			  Expect(em.sigma).To(Equal, mat)
		  })
		})
	})
}
