package obj_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) toString1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	e.Name = "yinggaimeigai"
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//为了避免内存拷贝
func (e *Employee) toString2() string {
	fmt.Print(e.Name)
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	e.Name = "gaile"
	return fmt.Sprintf("Id:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestObjDefine(t *testing.T) {
	//实例创建和初始化
	e := new(Employee)
	e.Id = "1"
	e.Name = "a"
	e.Age = 10
	fmt.Printf("Address is %p\n", e)
	fmt.Printf("Address is %p\n", &e.Id)

	var a Employee
	fmt.Printf("size is %d\n", unsafe.Sizeof(a))
	fmt.Printf("Address is %p\n", &a)
	fmt.Printf("Address is %p\n", &a.Id)
}
