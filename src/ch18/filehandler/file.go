package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("请输入待遍历目录")
	var path string
	fmt.Scan(&path)
	result := make(map[string]int)
	scandir(path, result)
	for k, v := range result {
		fmt.Printf("word:%s,count:%d\n", k, v)
	}
}

func scandir(dirPath string, result map[string]int) {
	dirFile, err := os.OpenFile(dirPath, os.O_RDONLY, os.ModeDir)
	defer dirFile.Close()
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	subFile, err := dirFile.ReadDir(-1)
	for _, v := range subFile {
		if v.IsDir() {
			scandir(dirPath+"/"+v.Name(), result)
		} else {
			readFile(dirPath+"/"+v.Name(), result)
		}
	}
}

func readFile(filePath string, result map[string]int) {
	if strings.HasSuffix(filePath, "a.txt") {
		fmt.Printf("filePath:%s\n", filePath)
		txt, err := os.OpenFile(filePath, os.O_RDONLY, 6)
		defer txt.Close()
		if err != nil {
			fmt.Println("OpenFile txt err:", err)
		}
		reader := bufio.NewReader(txt)
		for {
			bytes, err := reader.ReadBytes('\n')
			if err != nil && err == io.EOF {
				break
			}
			line := string(bytes)
			fields := strings.Fields(line)
			for _, word := range fields {
				if v, ok := result[word]; ok {
					result[word] = v + 1
				} else {
					result[word] = 1
				}
			}
		}

	}
}
