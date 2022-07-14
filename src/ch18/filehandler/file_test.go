package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	f, err := os.Create("/Users/wangbin/Documents/gofile.txt")
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}
	fmt.Println("f:", f)
	defer f.Close()
	fmt.Println("success")
}

func TestOpen(t *testing.T) {
	f, err := os.Open("/Users/wangbin/Documents/gofile.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("f:", f)
	_, err1 := f.WriteString("test write content to file ")
	if err1 != nil {
		fmt.Println("write file err:", err1)
		return
	}
	fmt.Println("success")
}

func TestOpenFile(t *testing.T) {
	f, err := os.OpenFile("/Users/wangbin/Documents/gofile.txt", os.O_RDWR, 6)
	defer f.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	fmt.Println("f:", f)
	_, err1 := f.WriteString("test write content to file ")
	if err1 != nil {
		fmt.Println("write file err:", err1)
		return
	}
	fmt.Println("success")
}

func TestWriteFile(t *testing.T) {
	f, err := os.OpenFile("/Users/wangbin/Documents/gofile.txt", os.O_RDWR, 6)
	defer f.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	//fmt.Println("f:", f)
	//_, err1 := f.WriteString("**")
	//if err1 != nil {
	//	fmt.Println("write file err:", err1)
	//	return
	//}
	seek, err := f.Seek(-1, io.SeekStart)
	at, err := f.WriteAt([]byte("*"), seek)
	fmt.Println("success:", at)
}

func TestReadFile(t *testing.T) {
	f, err := os.OpenFile("/Users/wangbin/Documents/gofile.txt", os.O_RDWR, 6)
	defer f.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	reader := bufio.NewReader(f)
	for {
		//按行读取
		buf, err := reader.ReadBytes('\n')
		if err == io.EOF { //文件结束标记
			break
		} else if err != nil {
			fmt.Println("read file err:", err)
		}
		fmt.Println(string(buf))
	}

}

func TestDir(t *testing.T) {
	dir, err := ioutil.ReadDir("/Users/wangbin/Documents/")
	if err != nil {

	}
	for i, v := range dir {
		fmt.Println("i:", i, "v:", v.Name())
	}

}
