package internal

import (
	"io/ioutil"
	"log"
)

func ScanFolder(argPath string) ([]string, error) {
	var fileList = []string{}

	files, error := ioutil.ReadDir(argPath)
	if error != nil {
		log.Fatal(error)
		return []string{}, error
	}

	for _, file := range files {
		if !file.IsDir() && ExtensionValidator(file.Name(), AccessExtension) {
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
