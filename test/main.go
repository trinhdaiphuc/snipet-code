package main

import (
	"fmt"
	"reflect"
)

func CheckType(a interface{}) string {
	return reflect.TypeOf(a).Kind().String()
}

func Reverse(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}
	result := make([]int, n)

	for i := 0; i < n/2+1; i++ {
		result[i] = a[n-i-1]
		result[n-i-1] = a[i]
	}

	return result
}

func main() {
	var A, B map[int]bool

	A = make(map[int]bool)
	B = make(map[int]bool)

	A[1] = true
	A[2] = false
	A[100] = true

	for k, v := range A {
		B[k] = v
	}

	fmt.Println("Map B", B)

	var (
		num            = 5
		str            = "Hello"
		numInt32 int32 = 10
	)

	fmt.Println("Type of num", CheckType(num))
	fmt.Println("Type of numInt32", CheckType(numInt32))
	fmt.Println("Type of str", CheckType(str))

	var arr = make([]int, 6)
	for i := 0; i < 6; i++ {
		arr[i] = i + 1
	}

	fmt.Println("Reverse array", Reverse(arr))
}
