package Panic_And_Recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "resolved")
		}
	}()
	fmt.Println("Program start.")
	panic(errors.New("an unknown error"))
}
