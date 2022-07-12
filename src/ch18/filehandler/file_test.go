package filehandler

import (
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
