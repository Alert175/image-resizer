package internal

import (
	"io/ioutil"
	"log"
	"os"
)

// Рекурсивное сканирование папки и получение ссылок на файлы
func ScanFolder(argPath string) ([]string, error) {
	var fileList = []string{}

	files, error := ioutil.ReadDir(argPath)
	if error != nil {
		log.Fatal(error)
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
