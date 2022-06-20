package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

// Рекурсивное сканирование папки и получение ссылок на файлы
func ScanFolder(argPath string) ([]string, error) {
	var fileList = []string{}

	files, error := ioutil.ReadDir(argPath)
	if error != nil {
		fmt.Println(error)
		return []string{}, error
	}

	for _, file := range files {
		if !file.IsDir() && ExtensionValidator(file.Name(), AccessExtensions) {
			filePath := argPath + "/" + file.Name()
			fileList = append(fileList, filePath)
		}
		if file.IsDir() {
			internalFileList, internalError := ScanFolder(argPath + "/" + file.Name())
			if internalError != nil {
				return fileList, error
			}
			fileList = append(fileList, internalFileList...)
		}

	}
	return fileList, nil
}

// Проверка на существование папки
func CheckFolder(argPath string) error {
	_, error := ioutil.ReadDir(argPath)
	if error != nil {
		return error
	}
	return nil
}

// создание вложенных папок
func CreateFolder(argPathFolder string) error {
	err := os.MkdirAll(argPathFolder, 0777)
	if err != nil {
		return err
	}
	return nil
}

// создание вложенных папок
func RemoveFolder(argPathFolder string) error {
	err := os.RemoveAll(argPathFolder)
	if err != nil {
		return err
	}
	return nil
}

// открыть папку через проводник ОС
func OpenWidthExplorer(argPath string) error {
	if runtime.GOOS == "linux" {
		if err := exec.Command("xdg-open", argPath).Start(); err != nil {
			return err
		}
		return nil
	}
	if runtime.GOOS == "darwin" {
		if err := exec.Command("open", argPath).Start(); err != nil {
			return err
		}
		return nil
	}
	if runtime.GOOS == "windows" {
		if err := exec.Command("explorer", argPath).Start(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("not found os name")
}
