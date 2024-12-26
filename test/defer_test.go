package test

import (
	"errors"
	"fmt"
	"testing"
)

func ReturnError() (err error) {
	err = nil
	defer func() {
		err := errors.New("test failed")
		fmt.Println(err)
	}()
	return err
}

func TestReturnError(t *testing.T) {
	err := ReturnError()
	fmt.Println(err)
}
