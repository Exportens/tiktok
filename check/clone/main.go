package main

import (
	"log"
	"os"
	"os/exec"
)
/////////////////////////////
	// 设置克隆目录
	cloneDir := "/home/ubuntu/go/src/github.com/cloudwego/api_gateway/idl"

	// 设置存储库URL
	repoURL := "https://github.com/example/repo.git"

	err := GitClone(cloneDir, repoURL)
	if err != nil {
		log.Fatal(err)
	}
/////////////////////////////

func GitClone(cloneDir, repoURL string) error {
	// 删除原文件夹
	err := os.RemoveAll(cloneDir)
	if err != nil {
		return err
	}

	// 执行git clone命令
	cmd := exec.Command("git", "clone", repoURL, cloneDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	log.Println("Git clone completed successfully.")
	return nil
}
