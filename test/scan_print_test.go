package test

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	var (
		s string
		i int
	)
	fmt.Scanf("%s%d", &s, &i)
	fmt.Printf("%s %d", s, i)
}
