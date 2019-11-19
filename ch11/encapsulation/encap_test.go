package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.Name, e.Age)
}

// 此方法会复制一个新的结构体然后再做打印的操作
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
// }

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Jayson", 35}
	e1 := Employee{Name: "Mike", Age: 24}
	e2 := new(Employee)
	e2.ID = "3"
	e2.Name = "Jacky"
	e2.Age = 30

	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	t.Logf("e is %T ", e)
	t.Logf("e2 if %T", e2)
}

func TestStructOperation(t *testing.T) {
	e := Employee{"0", "Jayson", 35}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
