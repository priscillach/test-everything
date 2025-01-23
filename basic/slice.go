package basic

import "fmt"

func NewSlice() {
	ps := new([]string)
	*ps = append(*ps, "Hello, World!")
}

func StringEdit() {
	s := "Hello, World!"
	// s[0] = 'h' immutable string
	s2 := s[1:5]
	fmt.Println(s2)
}
