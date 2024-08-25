package fs

import (
	"os"
)

// GetLastModifyTime 获取最后修改时间
func GetLastModifyTime(path string) (ts int64, err error) {
	var (
		f  *os.File
		fi os.FileInfo
	)
	if f, err = os.Open(path); err == nil {
		if fi, err = f.Stat(); err == nil {
			ts = fi.ModTime().Unix()
			_ = f.Close()
		}
	}
	return
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// SaveFile 将给定文本保存为文件
func SaveFile(fileName string, fileContent string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.WriteString(fileContent)
	if err != nil {
		return err
	}
	_ = f.Close()
	return nil
}

// CreateDirIfNotExists 目录不存在时创建目录
func CreateDirIfNotExists(path string) error {
	if !Exists(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
