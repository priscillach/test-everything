package main

import "fmt"

// 情况一
func case1() {
	slice := make([]int, 0, 4)
	slice = append(slice, 1, 2, 3)
	case11(slice)
	// slice的len没有变，print只有3个
	fmt.Println(slice)
}

func case11(slice []int) {
	slice = append(slice, 4)
}

// 情况二
func case2() {
	slice := make([]int, 0, 4)
	slice = append(slice, 1, 2, 3)
	case21(slice)
	fmt.Println(slice)
}

func case21(slice []int) {
	slice = append(slice, 4)
	slice[0] = 10
}

// 情况三
func case3() {
	slice := make([]int, 0, 3)
	slice = append(slice, 1, 2, 3)
	case31(slice)
	fmt.Println(slice)
}

func case31(slice []int) {
	slice = append(slice, 4)
	slice[0] = 10
}

func main() {
	case1()
	case2()
	case3()

	slice1 := make([]int, 0, 4)
	slice1 = append(slice1, 1, 2, 3)

	slice2 := slice1[:len(slice1)-1]
	slice2 = append(slice2, 11, 12, 13, 14, 15)
	slice2[0] = 10

	fmt.Println(slice1)
	fmt.Println(slice2)

}
