package basic

import (
	"fmt"
	"reflect"
	"unsafe"
)

func unsafeEditString() {
	// 使用 []byte 创建字符串，确保底层数据可写
	origBytes := []byte("hello")
	s := *(*string)(unsafe.Pointer(&origBytes))
	s2 := unsafe.String(&origBytes[0], 100)
	fmt.Println("Before s2:", s2) // 输出: hello
	fmt.Println("Before: s", s)   // 输出: hello
	origBytes[3] = 'Y'
	fmt.Println("Before s2:", s2) // 输出: hello
	fmt.Println("Before s:", s)   // 输出: hello

	// 利用 reflect.StringHeader 访问字符串的底层数据
	//s = "12345"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// 通过构造 []byte 切片访问同一底层数据
	bs := unsafe.Slice((*byte)(unsafe.Pointer(sh.Data)), sh.Len)
	bs[1] = 'a'
	fmt.Println("After s2:", s2) // 输出: hallo
	fmt.Println("After s:", s)   // 输出: hallo
}
