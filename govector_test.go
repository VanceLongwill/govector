package govector_test

import (
	"fmt"

	"github.com/vancelongwill/govector"
)

func ExampleSum() {
	a := govector.NewVector(1, 2, 3, 4)
	b := govector.NewVector(4, 5, 6, 7)
	fmt.Println(govector.Sum(a, b))
	// Output: [5 7 9 11]
}

func ExampleSubtract() {
	a := govector.NewVector(1, 2, 3)
	b := govector.NewVector(4, 5, 6)
	fmt.Println(govector.Subtract(a, b))
	// Output: [-3 -3 -3]
}

func ExampleScale() {
	a := govector.NewVector(1, 2, 3)
	fmt.Println(govector.Scale(a, 10))
	// Output: [10 20 30]
}

func ExampleDotProduct() {
	a := govector.NewVector(1, 2, 3)
	b := govector.NewVector(4, 5, 6)
	c := govector.NewVector(7, 8, 9)
	if product, err := govector.DotProduct(a, b, c); err == nil {
		fmt.Println(product)
	} else {
		fmt.Println(err)
	}
	// Output: 270
}

func ExampleCrossProduct() {
	a := govector.NewVector(1, 2, 3)
	b := govector.NewVector(4, 5, 6)
	if product, err := govector.CrossProduct(a, b); err == nil {
		fmt.Println(product)
	} else {
		fmt.Println(err)
	}
	// Output: [-3 6 -3]
}

func ExampleVector_Magnitude() {
	a := govector.NewVector(2, 2, 2, 2)
	fmt.Println(a.Magnitude())
	// Output: 4
}
