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
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	e.Name="yinggaimeigai"
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//为了避免内存拷贝
func (e *Employee) toString2() string {
	fmt.Print(e.Name)
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	e.Name="gaile"
	return fmt.Sprintf("Id:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}


func TestObjDefine(t *testing.T) {
	//实例创建和初始化
	e := Employee{"1", "Job", 10}
	t.Log(&e)
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	t.Log(e.toString1())
	t.Log(e.toString2())

	e1 := Employee{Id: "2", Name: "May", Age: 15}
	e2 := new(Employee)
	e2.Id = "5"
	e2.Name = "Hob"
	e2.Age = 23
	e3 := &Employee{}
	e.Id = "3"
	t.Log(e.Id)
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Log(e3)
	t.Logf("e1 is %T", e1)
	t.Logf("e2 is %T", e2)
	t.Logf("e3 is %T", e3)

}
