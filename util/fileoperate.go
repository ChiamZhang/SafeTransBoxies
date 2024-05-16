package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(FilePath string) (int, []byte, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("打开文件时出错：", err)
		return 0, nil, err
	}

	defer file.Close()

	// 获取文件信息，以确定文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息时出错：", err)
		return 0, nil, err
	}

	// 根据文件大小动态分配读取数据的缓冲区大小
	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	// 读取文件内容并显示在控制台
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("读取文件时出错：", err)
		return 0, nil, err
	}
	fmt.Println("文件内容：", string(data[:count]))
	return count, data, nil

}
func WriteToFile(filePath string, content []byte) error {
	// 将内容写入文件
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return err
	}
	return nil

}
