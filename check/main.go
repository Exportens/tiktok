package main

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"net/http"
)

func main() {
	path := "/home/ubuntu/go/src/github.com/cloudwego/api_gateway/check/file.txt" // 要检查的文件路径
	url := "http://localhost:8080/updated"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	var previousModTime time.Time // 保存上一次的修改时间
	fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("无法获取文件信息:", err)
			return
		}

		modTime := fileInfo.ModTime()

	previousModTime = modTime

	for {
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("无法获取文件信息:", err)
			return
		}

		modTime := fileInfo.ModTime()
		

		if !modTime.Equal(previousModTime) {
			fmt.Println("文件已更新")
			// 执行相应操作，例如重新加载文件
			// ...
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("发送请求失败:", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("读取响应失败:", err)
				return
			}
			// 打印响应内容
			fmt.Println("响应内容:", string(body))

			previousModTime = modTime
		}
		time.Sleep(time.Second) // 每秒钟检查一次更新
	}
}

