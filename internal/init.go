package internal

import (
	"fmt"
	"log"
)

// список ссылок на пути файлов
var FilePathList = []string{}

// список разрешенных расширений файлов
var AccessExtension = []string{"jpg"}

// Получение данных об содержании каталога input
func InitApp() {
	fmt.Println("Canninng...")
	fileList, errorFileList := ScanFolder("./input")
	if errorFileList != nil {
		log.Fatal(errorFileList)
	}
	FilePathList = append(FilePathList, fileList...)
	fmt.Println("Найдено файлов:", len(FilePathList))
	JpegComplession(FilePathList[0])
	resultFileList, errorFileList := ScanFolder("./output")
	if errorFileList != nil {
		log.Fatal(errorFileList)
	}
	fmt.Println("Создано файлов:", len(resultFileList))
}
