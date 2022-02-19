package main

// CS 311 Spring 2022
// Code with a simple main() to test your hw1 functions
import (
	"fmt"
	"math"
)

/*This function computes the value of polynomial equation for given co-efficients & x and returns the value
@parameter a float64
@parameter b float64
@parameter c float64
@parameter x float64
return float64 - value of polynomial equation
@author Rashmi Patel 25/01/22
*/
func compPoly(a, b, c, x float64) float64 {
	return ((a * (x * x)) + (b * x) + c)
}

/* This function computes two roots of quadratic formula for given a,b,c and return both the roots
@parameter a float64
@parameter b float64
@parameter c float64
return (float64, float64) - two roots of quadratic formala
@author Rashmi Patel 25/01/22
*/
func calcRoots(a, b, c float64) (float64, float64) {
	var root1, root2 float64
	if a != 0 {
		var determinant float64 = (b * b) - (4 * a * c)

		if determinant > 0 {
			root1 = ((-b) - math.Sqrt(determinant)) / (2 * a)
			root2 = ((-b) + math.Sqrt(determinant)) / (2 * a)
		} else {
			root1 = math.NaN()
			root2 = math.NaN()
		}
	} else {
		fmt.Println("Value of a is 0. Divide by zero error can occure.")
	}

	return root1, root2
}

/* This function checks if the given slice is in strictly descending order or not. Based on boolean value is returned.
@parameter []string - slice of strings
return bool - true if slice is in decending order else false
@author Rashmi Patel 25/01/22
*/
func lastToFirst(strs []string) bool {
	flag := false

	for i := 0; i < (len(strs) - 1); i++ {
		if strs[i] > strs[i+1] {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}

func main() {

	fmt.Println("--------------------------------compPoly - the value of polynomial equation--------------------------")
	// Should see 7.25
	fmt.Println(compPoly(1, 2, -4, 2.5))

	fmt.Println("-----------------------------calcRoots - two roots of quadratic formula------------------------------")
	// Should see 0.5 5
	fmt.Println(calcRoots(2, -11, 5))
	// Should see NaN NaN
	fmt.Println(calcRoots(3.14, 1, 4))

	fmt.Println("-----------------------lastToFirst - Checks if list is in strictly descending order------------------")
	strs1 := []string{"the", "lazy", "dog"}
	strs2 := []string{"the", "quick", "brown", "fox"}
	// Should see true
	fmt.Println(lastToFirst(strs1))
	// Should see false
	fmt.Println(lastToFirst(strs2))
}
