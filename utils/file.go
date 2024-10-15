package utils

import (
	"bufio"
	"fmt"
	"io"
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

	return err == nil
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
			log.Println(err)
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

func AppendToFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func CopyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer func(src *os.File) {
		_ = src.Close()
	}(src)

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		_, err = dst.Write(buf[:n])
		if err != nil {
			return err
		}
	}

	return nil
}
