package main //包，表明代码所在的模块

//应用程序入口
//必须是main包  package main
//必须是main方法 func main()
//文件名不一定是main.go

import (
	"flag"
	"fmt"
) //引入代码依赖

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

//退出返回值
//与其他主要编程语言的差异
//Go中main函数不支持任何返回值
//通过os.Exit来返回状态

//获取命令行参数
//与其他主要编程语言的差异
//main函数不支持传入参数
//func main(arg []string)
//在程序中直接通过os.Args获取命令行参数

//The master has failed more times than the beginner has tried
//功能实现
func main() {
	flag.Parse()
	fmt.Printf("Hello World:%s\n", name)
	// if len(os.Args) > 1 {
	// 	fmt.Println("Hello World", os.Args[1])
	// }
	// os.Exit(-1)
}
