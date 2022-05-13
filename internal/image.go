package internal

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func ImageResize(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	imageExtension := strings.Split(filePath, ".")[len(strings.Split(filePath, "."))-1]
	imageName := strings.Split(strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1], ".")[0]
	imagePath := strings.ReplaceAll(strings.ReplaceAll(filePath, "/"+imageName+"."+imageExtension, ""), "./input", "./output")

	if !IsJpg(imageExtension) && !IsPng(imageExtension) {
		log.Fatal("Не найдено валидное расширение файла")
		return "", err
	}

	var imageContent image.Image

	if IsJpg(imageExtension) {
		imgC, err := jpeg.Decode(file)
		if err != nil {
			return "", err
		} else {
			imageContent = imgC
		}
	}
	if IsPng(imageExtension) {
		imgC, err := png.Decode(file)
		if err != nil {
			return "", err
		} else {
			imageContent = imgC
		}
	}

	file.Close()

	compressImage := resize.Resize(500, 0, imageContent, resize.NearestNeighbor)
	fmt.Println(imagePath)
	if err := CheckFolder(imagePath); err != nil {

	} else {

	}
	compressFile, err := os.Create("./output/" + imageName + "." + imageExtension)

	if err != nil {
		log.Fatal(err)
	}
	defer compressFile.Close()

	if IsJpg(imageExtension) {
		jpeg.Encode(compressFile, compressImage, nil)
	}
	if IsPng(imageExtension) {
		png.Encode(compressFile, compressImage)
	}

	return "./output/" + imageName + "." + imageExtension, nil
}

// проверка, на то что изображение формата jpg
func IsJpg(argExtension string) bool {
	for _, element := range []string{"jpeg", "jpg"} {
		if argExtension == element {
			return true
		}
	}
	return false
}

// проверка, на то что изображение формата png
func IsPng(argExtension string) bool {
	for _, element := range []string{"png"} {
		if argExtension == element {
			return true
		}
	}
	return false
}

func JpegComplession(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	compressImage := resize.Resize(500, 0, img, resize.Lanczos3)

	compressFile, err := os.Create("./output/test_imge.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer compressFile.Close()

	jpeg.Encode(compressFile, compressImage, nil)
}

func PngComplession(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	compressImage := resize.Resize(300, 0, img, resize.Bicubic)

	compressFile, err := os.Create("./output/test_imge.png")
	if err != nil {
		log.Fatal(err)
	}
	defer compressFile.Close()

	png.Encode(compressFile, compressImage)
}
