package Empty_Interface

import (
	"fmt"
	"testing"
)

func CheckType(p interface{}) {
	//if i, j := p.(int); j {
	//	fmt.Println("Integer:\t", i)
	//}
	//if i, j := p.(string); j {
	//	fmt.Println("String:\t", i)
	//}
	//fmt.Println("UNKNOWN")
	switch data := p.(type) {
	case int:
		fmt.Println("Integer:\t", data)
	case string:
		fmt.Println("String:\t", data)
	default:
		fmt.Println("UNKNOWN")
	}
}
func TestEmptyInterface(t *testing.T) {
	CheckType(114514)
	CheckType("Cyka Blyat")
	CheckType('0')
}
