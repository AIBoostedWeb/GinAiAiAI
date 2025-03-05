package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRoot(markerFile string) (string, error) {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取工作目录失败: %w", err)
	}

	// 向上查找标记文件（如 go.mod）
	dir := wd
	for {
		if _, err := os.Stat(filepath.Join(dir, markerFile)); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir { // 到达根目录
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("未找到标记文件 %s", markerFile)
}
