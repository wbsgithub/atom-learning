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

func TestStructInit(t *testing.T) {
	//var emp Employee = Employee{"a", "a1", 1}
	//fmt.Println(emp)
	//fmt.Println(unsafe.Pointer(&emp))
	//modifyEmp(emp)
	emp2 := new(Employee)
	emp2.Id = "1"
	emp2.Name = "a2"
	emp2.Age = 10
	fmt.Println(emp2)
	fmt.Println(&emp2)
	modifyEmpPointer(emp2)
	fmt.Println(emp2)
	fmt.Println(&emp2)

	var emp3 Employee
	fmt.Println(emp3)
	fmt.Println(&emp3)
	emp3.Id = "2"
	emp3.Name = "a3"
	emp3.Age = 14
	modifyEmpPointer(&emp3)
	fmt.Println(&emp3)

}

func modifyEmpPointer(emp *Employee) {
	fmt.Println(&emp) //&emp表示复制的指针地址值存放的地址
	emp.Name = "b2"
}

func modifyEmp(emp Employee) {
	emp.Name = "b"
	fmt.Println(emp)
	fmt.Println(unsafe.Pointer(&emp))
}
