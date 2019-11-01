package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetFileData(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// 出现错误则退出
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	// 默认当前路径
	path := `./`
	// 获取 当前文件的前 5 行
	var cmake_file = "CMakeLists.txt"
	filedata := GetFileData(cmake_file)
	list := strings.Split(filedata, "\n")
	list = list[:5]
	// 获取当前路径的文件列表
	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	//// 遍历文件列表向文件写入内容
	f, err := os.Create(cmake_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// 写入内容到文件
	for _, v := range list {
		f.WriteString(v + "\n")
	}
	for _, file := range fileInfoList {
		index := strings.LastIndex(file.Name(), ".cpp")
		if index != -1 {
			f.WriteString("add_executable(" + file.Name() + " " + file.Name() + ")\n")
			fmt.Println(file.Name())
		}
	}
	fmt.Println("已完成")
}
