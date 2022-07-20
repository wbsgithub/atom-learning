package main

import (
	"bufio"
	"fmt"
	"io"
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
		fmt.Println(string(buf[:]))
	}

}

func TestFileCopy(t *testing.T) {
	source, err := os.OpenFile("/Users/wangbin/Documents/区块链课程/2019年Go语言与区块链在线就业班/第一、二阶段/day1/视频/01.指针地址个变量空间.avi", os.O_RDWR, 6)
	defer source.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	target, err := os.Create("/Users/wangbin/Documents/区块链课程/2019年Go语言与区块链在线就业班/第一、二阶段/day1/视频/01.指针地址个变量空间1.avi")
	defer target.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}
		if n == 0 {
			fmt.Println("文件复制完毕")
			break
		}
		tmp := buf[:n]
		target.Write(tmp)
	}

}
