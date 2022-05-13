package internal

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"github.com/nfnt/resize"
)

func ImageResize(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	imageExtension := strings.Split(filePath, ".")[len(strings.Split(filePath, ".")) - 1]
	imageName := strings.Split(strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1], ".")[0]
	isJpg =: strings.Contains()
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

	jpeg.Encode(compressFile, compressImage, nil)
}
