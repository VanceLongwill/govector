package govector

import (
	"errors"
	"math"
)

var (
	// ErrOutOfRange is returned when trying to access an component under an index beyond the defined dimensions of the vector
	ErrOutOfRange = errors.New("Vector component index out of range")
	// ErrVectorNot3D is returned when a 3D vector is expected
	ErrVectorNot3D = errors.New("Vector is not 3D")
)

// Vector represents a mathematical vector with an abitrary number of components (dimensions)
type Vector []float64

// At gets the component value at a specific index, returns an error if it can't be retrieved
func (v Vector) At(n int) (float64, error) {
	if n < len(v) {
		return v[n], nil
	}
	return 0, ErrOutOfRange
}

// Count returns the number of components (dimensions) in the vector
func (v Vector) Count() int {
	return len(v)
}

// Magnitude calculates the magnitude (length) of the vector for any given number of components
func (v Vector) Magnitude() float64 {
	var sumAllSquared float64
	for _, component := range v {
		sumAllSquared += math.Pow(component, 2)
	}
	return math.Pow(sumAllSquared, 0.5)
}

// NewVector creates a new vector from a slice of components
func NewVector(components ...float64) Vector {
	return components
}

// getMaxComponentCount find the greatest number of components present in any of the given vectors
func getMaxComponentCount(vectors ...Vector) int {
	// find the vector with the greatest number of components
	var componentsCount int = 0
	for _, v := range vectors {
		if c := v.Count(); c > componentsCount {
			componentsCount = c
		}
	}
	return componentsCount
}

// Sum an abitrary number of vectors
func Sum(vectors ...Vector) Vector {
	components := make([]float64, getMaxComponentCount(vectors...))
	for i := range components {
		for _, v := range vectors {
			if component, err := v.At(i); err != nil {
				components[i] += component
			}
		}
	}
	return NewVector(components...)
}

// Subract an abitrary number of vectors
func Subract(vectors ...Vector) Vector {
	components := make([]float64, getMaxComponentCount(vectors...))
	for i := range components {
		for _, v := range vectors {
			if component, err := v.At(i); err != nil {
				components[i] -= component
			}
		}
	}
	return NewVector(components...)
}

// Scale performs a scalar multiplication on a vector, returning a new vector
func Scale(v Vector, scalar float64) Vector {
	components := make([]float64, v.Count())
	for i := range components {
		// no need to check error since we have set the components slice size as the number of components
		c, _ := v.At(i)
		components[i] = c * scalar
	}
	return NewVector(components...)
}

// DotProduct calculates the dot product scalar value of an abitrary number of vectors
func DotProduct(vectors ...Vector) (float64, error) {
	var sumOfProducts float64
	for i := 0; i < getMaxComponentCount(vectors...); i++ {
		var productOfComponents float64
		for _, v := range vectors {
			if component, err := v.At(i); err != nil {
				productOfComponents *= component
			} else {
				return 0, err
			}
		}
		sumOfProducts += productOfComponents
	}
	return sumOfProducts, nil
}

// is3DVector determines if a vector is in 3D space
func is3DVector(v Vector) bool {
	return v.Count() == 3
}

// CrossProduct calculates the cross product of two 3D vectors, returning a new vector
func CrossProduct(a Vector, b Vector) (Vector, error) {
	if !is3DVector(a) || !is3DVector(b) {
		return nil, ErrVectorNot3D
	}
	// set up values for the matrix
	a1, _ := a.At(0)
	a2, _ := a.At(1)
	a3, _ := a.At(2)

	b1, _ := b.At(0)
	b2, _ := b.At(1)
	b3, _ := b.At(2)

	components := make([]float64, 3)
	components[0] = (a2 * b3) - (a3 * b2)
	components[1] = (a3 * b1) - (a1 * b3)
	components[2] = (a1 * b2) - (a2 * b1)

	return NewVector(components...), nil
}
