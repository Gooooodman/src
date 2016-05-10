package main

import "fmt"

//传地址
// func main() {
// 	array := [3]float64{7.0, 8.5, 9.1}
// 	x := Sum(&array) // Note the explicit address-of operator
// 	// to pass a pointer to the array
// 	fmt.Printf("The sum of the array is: %f", x)
// }

// func Sum(a *[3]float64) (sum float64) {
// 	for _, v := range a { // derefencing *a to get back to the array is not necessary!
// 		sum += v
// 	}
// 	return
// }

//优选方案 传切片
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(sum(arr[:]))
}
