package util

import (
	"fmt"
	"os"
)

func MoveFileTo(src, dst string) {
	// 	src := "output.jpg"
	//     // 目标文件路径，这里假设/download文件夹已经存在
	//     dst := "/download/output.jpg"

	// 移动文件
	err := os.Rename(src, dst)
	if err != nil {
		fmt.Printf("Failed to move file: %s\n", err)
		return
	}

	fmt.Println("File moved successfully")
}
