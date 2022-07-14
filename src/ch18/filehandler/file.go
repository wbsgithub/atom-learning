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

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer file.Close()
	dir, err := file.ReadDir(-1)
	if err != nil {

	}
	result := make(map[string]int)
	for i, v := range dir {
		fmt.Println("i:", i, "v:", v.Name())
		if strings.HasSuffix(v.Name(), ".txt") {
			txt, err := os.OpenFile(path+"/"+v.Name(), os.O_RDWR, 6)
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
	for k, v := range result {
		fmt.Printf("word:%s,count:%d\n", k, v)
	}
}
