package internal

import (
	"fmt"
	"log"
	"os"
)

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

	var imageWidth int
	var resultFileCounter int

	fmt.Println("Найдено файлов:", len(fileList))
	fmt.Println("Укажите ширину изображений: ")
	fmt.Scan(&imageWidth)

	RemoveFolder("./output")
	CreateFolder("./output")

	for _, path := range fileList {
		_, err := ImageResize(path, imageWidth)
		if err != nil {
			log.Fatal(err)
		} else {
			resultFileCounter = resultFileCounter + 1
		}
	}

	_, resultErrorFileList := ScanFolder("./output")
	if resultErrorFileList != nil {
		log.Fatal(errorFileList)
	}

	fmt.Println("Создано файлов: ", resultFileCounter)
}
