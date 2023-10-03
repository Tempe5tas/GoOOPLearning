package structure

import (
	"fmt"
	"testing"
	"unsafe"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

//func (str Person) Print() string {	//Data space copy
//	fmt.Println("Function pointer:", unsafe.Pointer(&str.ID))
//	return fmt.Sprintf("ID: %d\tName: %s\tAge: %d", str.ID, str.Name, str.Age)
//}

func (str *Person) Print() string { //No copy
	fmt.Println("Function pointer: ", unsafe.Pointer(&str.ID))
	return fmt.Sprintf("ID: %d\tName: %s\tAge: %d", str.ID, str.Name, str.Age)
}

func TestStructDeclare(t *testing.T) {
	str := Person{114514, "Tadokoro", 24}
	str1 := new(Person) //Same as str1:=&Person{}
	str1.ID = 1919810
	str1.Name = "Tadokoro Koji"
	str1.Age = 19
	t.Log(str, str1)
	t.Logf("%T %T", str, str1)
}

func TestStructureFunc(t *testing.T) {
	str := &Person{114514, "Tadokoro", 24}
	fmt.Println("Structure pointer:", unsafe.Pointer(&str.ID))
	t.Log(str.Print())
}
