package internal

import (
	"fmt"
)

// список разрешенных расширений файлов
var AccessExtensions = []string{"jpg", "jpeg", "png"}

// Получение данных об содержании каталога input
func InitApp() {
	fmt.Println("Сканирование...")
	exitStr := ""
	fileList, errorFileList := ScanFolder("./input")
	if errorFileList != nil {
		fmt.Println("Ошибка чтения каталога input ", errorFileList)
	}
	if len(fileList) == 0 {
		fmt.Println("Файлы не обнаружены")
		fmt.Println("Введите что-либо для продолжения")
		fmt.Scan(&exitStr)
		return
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
			fmt.Println(err)
		} else {
			resultFileCounter = resultFileCounter + 1
			exitStr = exitStr + "-"
			fmt.Print("-")
		}
	}

	_, resultErrorFileList := ScanFolder("./output")
	if resultErrorFileList != nil {
		fmt.Println("Ошибка чтения каталога input ", resultErrorFileList)
	}
	fmt.Println("")
	fmt.Println("Создано файлов: ", resultFileCounter)
	fmt.Println("Введите что-либо для продолжения")
	fmt.Scan(&exitStr)
}
