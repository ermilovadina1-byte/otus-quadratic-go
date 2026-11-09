package quadratic

import (
	"math"
	"testing"
)

func TestNoRealRoots(t *testing.T) {
	roots, err := Solve(1, 0, 1)
	if err != nil {
		t.Fatalf("Solve(1, 0, 1) returned error: %v", err)
	}
	if len(roots) != 0 {
		t.Fatalf("Solve(1, 0, 1) = %v, want no real roots", roots)
	}
}

func TestTwoDistinctRoots(t *testing.T) {
	roots, err := Solve(1, 0, -1)
	if err != nil {
		t.Fatalf("Solve(1, 0, -1) returned error: %v", err)
	}
	if len(roots) != 2 || roots[0] != 1 || roots[1] != -1 {
		t.Fatalf("Solve(1, 0, -1) = %v, want [1 -1]", roots)
	}
}

func TestDoubleRoot(t *testing.T) {
	roots, err := Solve(1, 2, 1)
	if err != nil {
		t.Fatalf("Solve(1, 2, 1) returned error: %v", err)
	}
	if len(roots) != 2 || roots[0] != -1 || roots[1] != -1 {
		t.Fatalf("Solve(1, 2, 1) = %v, want [-1 -1]", roots)
	}
}

func TestNearZeroAIsInvalid(t *testing.T) {
	for _, a := range []float64{0, 1e-7, -1e-7} {
		_, err := Solve(a, 1, 1)
		if err == nil {
			t.Errorf("Solve(%g, 1, 1) returned no error, want invalid a", a)
		}
	}
}

func TestSmallNonzeroDiscriminantCountsAsDoubleRoot(t *testing.T) {
	a, b, c := 1.0, 2.0, 1.0-1e-7
	discriminant := b*b - 4*a*c
	if discriminant <= 0 || discriminant >= Epsilon {
		t.Fatalf("test input must have 0 < discriminant < epsilon, got %g", discriminant)
	}
	roots, err := Solve(a, b, c)
	if err != nil {
		t.Fatalf("Solve(%g, %g, %g) returned error: %v", a, b, c, err)
	}
	if len(roots) != 2 || roots[0] != -1 || roots[1] != -1 {
		t.Fatalf("Solve(%g, %g, %g) = %v, want [-1 -1]", a, b, c, roots)
	}
}

func TestNonFiniteCoefficientIsInvalid(t *testing.T) {
	special := []struct {
		name  string
		value float64
	}{
		{"NaN", math.NaN()},
		{"positive infinity", math.Inf(1)},
		{"negative infinity", math.Inf(-1)},
	}
	for _, item := range special {
		for index, name := range []string{"a", "b", "c"} {
			coefficients := [3]float64{1, 2, 1}
			coefficients[index] = item.value
			_, err := Solve(coefficients[0], coefficients[1], coefficients[2])
			if err == nil {
				t.Errorf("%s as %s returned no error", item.name, name)
			}
		}
	}
}

func TestFiniteInputsMustNotProduceNonFiniteRoots(t *testing.T) {
	roots, err := Solve(1e308, 1e308, 1e308)
	if err == nil {
		for _, root := range roots {
			if math.IsNaN(root) || math.IsInf(root, 0) {
				t.Fatalf("finite coefficients produced non-finite root %g without error", root)
			}
		}
	}
}
