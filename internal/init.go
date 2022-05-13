package internal

import (
	"fmt"
	"log"
	"os"
)

// список ссылок на пути файлов
var FilePathList = []string{}

// список разрешенных расширений файлов
var AccessExtensions = []string{"jpg", "jpeg", "png"}

// Получение данных об содержании каталога input
func InitApp() {
	fmt.Println("Canninng...")
	fileList, errorFileList := ScanFolder("./input")
	if errorFileList != nil {
		log.Fatal(errorFileList)
	}
	if len(fileList) == 0 {
		fmt.Println("Файлы не обнаружены")
		os.Exit(0)
	}
	FilePathList = append(FilePathList, fileList...)
	fmt.Println("Найдено файлов:", len(FilePathList))
	// JpegComplession(FilePathList[0])
	// resultFileList, errorFileList := ScanFolder("./output")
	// if errorFileList != nil {
	// 	log.Fatal(errorFileList)
	// }
	// fmt.Println("Создано файлов:", len(resultFileList))
	ImageResize(FilePathList[0])
}
