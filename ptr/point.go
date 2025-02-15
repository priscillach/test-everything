package main

import (
	"fmt"
	"unsafe"
)

func TestPrt() {
	arr0 := make([]int, 0, 4)
	arr0 = append(arr0, 1, 2, 3, 4)
	arr1 := arr0
	arr1[0] = 5
	fmt.Println(arr0) // [5 2 3 4]
	fmt.Println(arr1) // [5 2 3 4]
	arr0 = append(arr0, 1, 2, 3, 4, 5)
	arr1[1] = 6
	fmt.Println(arr0) // [5 2 3 4 1 2 3 4 5]
	fmt.Println(arr1) // [5 6 3 4]

	arr3 := make([]int, 0, 4)
	arr3 = append(arr3, 1, 2, 3, 4)
	arr3P := &arr3

	// arr3和arr4就是同一个东西
	arr4 := (*[]int)(unsafe.Pointer(arr3P))
	*arr4 = append(*arr4, 5, 6, 7)
	fmt.Println(arr3)
	fmt.Println(*arr4)
	fmt.Printf("arr3 addr: %p\n", &arr3)
	fmt.Printf("arr3 data addr: %p\n", arr3)
	fmt.Printf("arr4 addr: %p\n", arr4)
	fmt.Printf("arr4 data addr: %p\n", *arr4)
	/*
		[1 2 3 4 5 6 7]
		[1 2 3 4 5 6 7]
		arr3 addr: 0x14000124060
		arr3 data addr: 0x1400011e040
		arr4 addr: 0x14000124060
		arr4 data addr: 0x1400011e040
	*/
	arr5 := make([]int, 0, 4)
	arr5 = append(arr5, 1, 2, 3, 4)
	arr6 := arr5[2:]
	arr6P := &arr6

	arr7 := (*[]int)(unsafe.Pointer(arr6P))
	*arr7 = append(*arr7, 5, 6, 7)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
}

func TestUnsafePtr() {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)
	eleSize := int(unsafe.Sizeof(a[0]))
	fmt.Println(eleSize)
	p9 := &a[9]
	up9 := unsafe.Pointer(p9)
	p3 := (*int)(unsafe.Add(up9, -6*eleSize))
	fmt.Println(*p3) // 3
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println(s)              // [9 0 11]
	fmt.Println(len(s), cap(s)) // 3 5

	t := unsafe.Slice((*int)(nil), 0)
	fmt.Println(t == nil) // true

	// 下面是两个不正确的调用。因为它们
	// 的返回结果引用了未知的内存块。
	_ = unsafe.Add(up9, 7*eleSize)
	_ = unsafe.Slice(p9, 8)
}

func main() {
	//TestPrt()
	TestUnsafePtr()
}
