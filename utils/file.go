package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %v", err)
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// 使用 filepath.ToSlash 替换路径中的反斜杠
	dir = filepath.ToSlash(dir)
	return dir, nil
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

func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		return 0, fmt.Errorf("failed to get file info: %v", err)
	}

	return fileInfo.Size(), nil
}

func MustReadFileString(filename string) string {
	content, _ := ReadFileString(filename)

	return content
}

func ReadFileString(filename string) (string, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return "", err
	}

	return string(fileBytes), nil
}

func WriteFileWithString(filename string, content string) error {
	return WriteFileWithReader(filename, strings.NewReader(content))
}

func WriteFileWithReader(filename string, reader io.Reader) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func() {
		ce := file.Close()
		if ce != nil {
			log.Printf("Failed to close file: %v", ce)
		}
	}()

	n, err := io.Copy(file, reader)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return fmt.Errorf("failed to write to file: %v", err)
	}

	log.Printf("Wrote %d bytes to file: %s", n, filename)
	return nil
}

func AppendToFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		ce := file.Close()
		if ce != nil {
			log.Printf("Failed to close file: %v", ce)
		}
	}()

	n, err := io.Copy(file, bytes.NewReader(data))
	if err != nil {
		return err
	}

	if n != int64(len(data)) {
		return fmt.Errorf("wrote %d bytes, expected %d bytes", n, len(data))
	}

	return nil
}

func CopyDir(srcDir, dstDir string) error {
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			err = CopyDir(sourcePath, destPath)
		} else {
			err = CopyFile(sourcePath, destPath)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func CopyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer func() {
		ce := src.Close()
		if ce != nil {
			log.Printf("Failed to close source file: %v", ce)
		}
	}()

	srcInfo, err := src.Stat()
	if err != nil {
		return err
	}

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer func() {
		ce := dst.Close()
		if ce != nil {
			log.Printf("Failed to close destination file: %v", ce)
		}
	}()

	n, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("Failed to copy file: %v", err)
		return err
	}

	if n != srcInfo.Size() {
		return fmt.Errorf("copied %d bytes, expected %d bytes", n, srcInfo.Size())
	}

	return nil
}

func CalcFileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer func() {
		ce := file.Close()
		if ce != nil {
			log.Printf("Failed to close file: %v", ce)
		}
	}()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Printf("Error hashing file: %v", err)
		return "", fmt.Errorf("failed to hash file: %v", err)
	}

	md5sum := hex.EncodeToString(hash.Sum(nil))
	log.Printf("MD5 checksum of the file(%s) is: %s", filename, md5sum)
	return md5sum, nil
}

func ListFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 如果是文件，则打印其路径
		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	return files, err
}

func ListDirectSubdirectories(dir string) ([]string, error) {
	var subdirectories []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subdirectory := filepath.Join(dir, file.Name())
			subdirectories = append(subdirectories, subdirectory)
		}
	}

	return subdirectories, nil
}

func ListDirectFiles(dir string) ([]string, error) {
	var directFiles []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			directFiles = append(directFiles, filepath.Join(dir, file.Name()))
		}
	}

	return directFiles, nil
}

func RenameSubdirectoriesBlankSpaceWithUnderline(dir string) error {
	subdirectories, err := ListDirectSubdirectories(dir)
	if err != nil {
		return err
	}
	if len(subdirectories) == 0 {
		return nil
	}

	var newSubdirectories []string
	for _, subdirectory := range subdirectories {
		if strings.Contains(subdirectory, " ") {
			newSubdirectory := strings.ReplaceAll(subdirectory, " ", "_")
			re := os.Rename(subdirectory, newSubdirectory)
			if re != nil {
				return re
			}
			newSubdirectories = append(newSubdirectories, newSubdirectory)
		} else {
			newSubdirectories = append(newSubdirectories, subdirectory)
		}
	}

	for _, subdirectory := range newSubdirectories {
		err = RenameSubdirectoriesBlankSpaceWithUnderline(subdirectory)
		if err != nil {
			return err
		}
	}

	return nil
}

func RenameSubFilesBlankSpaceWithUnderline(dir string) error {
	err := RenameSubdirectoriesBlankSpaceWithUnderline(dir)
	if err != nil {
		return err
	}

	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.Contains(path, " ") {
			re := os.Rename(path, strings.ReplaceAll(path, " ", "_"))
			if re != nil {
				return re
			}
		}

		return nil
	})
}

func GetFileBaseWithoutExt(filename string) string {
	fileBase := path.Base(filename)
	return fileBase[:strings.Index(fileBase, ".")]
}

func ReplaceFileExt(filename, toExt string) string {
	return filename[:strings.Index(filename, ".")] + toExt
}
