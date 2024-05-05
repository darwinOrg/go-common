package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func ExistsDir(dir string) bool {
	d, err := os.Stat(dir)
	if err != nil {
		log.Println(err)
		return false
	}
	return d.IsDir()
}

func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)

	return os.IsNotExist(err)
}

func GetFileSize(filename string) int64 {
	var size int64
	err := filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		size = f.Size()
		return nil
	})
	if err != nil {
		log.Println(err)
		return 0
	}
	return size
}

func WriteFileWithString(filename string, content string) error {
	file, err := os.Create(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
