package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	aResult1 := nums_A_after_predicate_1(2, 99)
	bResult1 := nums_B_after_predicate_1(aResult1)
	bResult2 := nums_B_after_predicate_2(bResult1)
	fmt.Printf("\n\n\n\n")
	sort.Ints(aResult1)
	for _, k := range aResult1 {
		if len(bResult2[k]) == 1 {
			fmt.Printf("Key:%d,    Value:%d\n", k, bResult2[k])
		}
	}
	fmt.Printf("\n\n\n\n")
}

func sumSplit(m int) map[int]int {
	var sumMap = make(map[int]int)
	// m > 2
	for i := 2; i < m/2+1; i++ {
		sumMap[i] = m - i
	}
	return sumMap
}

func factorize(n int) map[int]int {
	var productMap = make(map[int]int)
	temp := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < temp; i++ {
		if n%i == 0 {
			productMap[i] = n / i
		}
	}
	return productMap
}

func sum_split_has_unique_fac(m, mi, ma int) bool {
	sumMap := sumSplit(m)
	for p, q := range sumMap {
		productMap := factorize(p * q)
		for a, b := range productMap {
			if mi <= a && a <= ma && mi <= b && b <= ma {
				if len(productMap) == 1 {
					return true
				}
			}
		}
	}
	return false
}

func nums_A_after_predicate_1(mi, ma int) []int {
	var slice1 []int
	for m := 2 * mi; m < ma/2+10; m++ {
		result := sum_split_has_unique_fac(m, mi, ma)
		if !result {
			slice1 = append(slice1, m)
		}
	}
	return slice1
}

func nums_B_after_predicate_1(ms []int) map[int][]int {
	var theMap = make(map[int][]int)
	for _, x := range ms {
		sumMap := sumSplit(x)
		var productArray []int
		for a, b := range sumMap {
			productArray = append(productArray, a*b)
		}
		theMap[x] = productArray
	}

	return theMap
}

func nums_B_after_predicate_2(n1_to_prods map[int][]int) map[int][]int {
	var appear = make(map[int]int)
	for _, v := range n1_to_prods {
		for _, p := range v {
			appear[p]++
		}
	}
	// delete repeat element
	for p, a := range appear {
		if a > 1 {
			delete(appear, p)
		}
	}
	for k, v := range n1_to_prods {
		var slice1 []int
		for _, j := range v {
			if appear[j] == 1 {
				slice1 = append(slice1, j)
			}
		}
		n1_to_prods[k] = slice1
	}
	return n1_to_prods
}
