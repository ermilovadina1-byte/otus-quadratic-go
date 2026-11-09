package quadratic

import (
	"errors"
	"math"
)

const Epsilon = 1e-5

var ErrNotQuadratic = errors.New("coefficient a is too close to zero")
var ErrNonFiniteCoefficient = errors.New("coefficient is not finite")
var ErrUnrepresentableResult = errors.New("calculation is outside float64 range")

// Solve returns the real roots of a*x*x+b*x+c=0, including multiplicity.
// TODO: implement the cases described by the tests.
func Solve(a, b, c float64) ([]float64, error) {
	for _, coefficient := range []float64{a, b, c} {
		if math.IsNaN(coefficient) || math.IsInf(coefficient, 0) {
			return nil, ErrNonFiniteCoefficient
		}
	}
	if math.Abs(a) < Epsilon {
		return nil, ErrNotQuadratic
	}
	discriminant := b*b - 4*a*c
	if math.IsNaN(discriminant) || math.IsInf(discriminant, 0) {
		return nil, ErrUnrepresentableResult
	}
	if math.Abs(discriminant) < Epsilon {
		root := -b / (2 * a)
		return []float64{root, root}, nil
	}
	if discriminant < 0 {
		return []float64{}, nil
	}
	root := math.Sqrt(discriminant)
	return []float64{(-b + root) / (2 * a), (-b - root) / (2 * a)}, nil
}
