package commit

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"replacetools/src/config"
)

func applyUpdate(directory, updateName string) error {
	updatePath := filepath.Join("commit", updateName)
	return filepath.Walk(updatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // 忽略目录
		}
		// 构建目标文件路径
		relPath, err := filepath.Rel(updatePath, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(directory, relPath)
		return copyFile(path, targetPath)
	})
}

// 复制文件
func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, input, 0644)
	fmt.Println("已更新目标路径:" + dst)
	return err
}

func Commit(out config.Config, fileName string) {
	for _, dir := range out.Directories {
		applyUpdate(dir, fileName)
	}
	fmt.Println("已完成所有文件更新")
}
