package extention

import (
	"fmt"
	"testing"
	"unsafe"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Print("....")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct {
	//p *Pet
	//匿名嵌套
	Pet
}

//func (d *Dog) speak() {
//	d.p.speak()
//}
//
//func (d *Dog) speakTo(host string) {
//	d.p.speakTo(host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.speakTo("wangbin")
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

////第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) detail() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
//func (e *Employee) detail() string {
//	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("Id:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}

func TestEmp1(t *testing.T) {
	e := Employee{"1", "lilei", 10}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.detail())
	//e2 := &Employee{"1", "lilei", 10}
	//t.Log(e2.detail())

}
func TestEmp(t *testing.T) {
	e := Employee{"1", "lilei", 10}
	e1 := Employee{Id: "2", Name: "hle", Age: 30}
	e2 := new(Employee) //这里返回的引用/指针 ，相当于 e:=&Emplyee{}
	e2.Id = "3"         //通过实例的指针访问成员不需要使用->
	e2.Name = "adsf"
	e2.Age = 304

	t.Log(e)
	t.Log(e1)
	t.Log(e2)

	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)

}
