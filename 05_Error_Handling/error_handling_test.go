package Error_Handling

import (
	"errors"
	"testing"
)

func Fibonacci(i int) (int, error) {
	if i > 100 || i < 1 {
		return 0, errors.New("Input number should between [1,100]")
	}
	arr := []int{1, 1}
	for j := 2; j < i; j++ {
		arr = append(arr, arr[j-1]+arr[j-2])
	}
	return arr[i-1], nil
}
func TestFibonacci(t *testing.T) {
	if n, err := Fibonacci(-10); err != nil {
		t.Log(err)
	} else {
		t.Log(n)
	}
	t.Log(Fibonacci(10))
}
