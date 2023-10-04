package Unit_Test

import "testing"

func cube(n int) int {
	return n * n * n
	//return n*n*n+1
}

// Fail / Error: continue execution
// FailNow / Fatal: abort execution

// If execute test in terminal:
// go test -v -cover

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3, 4}
	expected := [...]int{1, 8, 27, 64}
	for i := 0; i < len(inputs); i++ {
		ret := cube(inputs[i])
		if ret != expected[i] {
			//t.Errorf("expected %d, returned %d.", expected[i], ret)
			t.Fatalf("expected %d, returned %d.", expected[i], ret)
		}
	}
}
