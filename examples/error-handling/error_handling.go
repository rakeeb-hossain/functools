// This showcases how we can handle errors using functools helpers, despite there being no explicit error
// return types by the functools functions. We make use of passing errors via closures.
//
// This is more explicit than it likely would be in practice just to showcase what safe handling capabilities
// this functools still enables. In practice, checking for errant input as a pre-step before dividing would likely
// be more effective.
package main

import (
	"errors"
	"github.com/rakeeb-hossain/functools"
	"log"
)

type fraction struct {
	dividend int
	divisor  int
}

func main() {
	fractions := []fraction{{5, 1}, {3, 6}, {2, 0}}

	// We handle errors by populating an error which we pass to our mapper function via a closure.
	// We also return a pointer to a float64 instead of a float64 itself, so we can handle nil types
	// in case we encounter an error.
	var err error
	safeDivide := func(f fraction) *float64 {
		if f.divisor == 0 {
			err = errors.New("cannot divide by 0")
			return nil
		}
		res := float64(f.dividend) / float64(f.divisor)
		return &res
	}
	rationalResults := functools.Map(fractions, safeDivide)
	if err != nil {
		log.Println(err)
	}

	// We can sum the safe rational results using a custom Reduce function
	res := functools.Reduce(rationalResults, 0.0, func(accum float64, n *float64) float64 {
		if n == nil {
			return accum
		}
		return accum + *n
	})

	log.Printf("the safe sum of fractions %v is %f\n", fractions, res)
}
