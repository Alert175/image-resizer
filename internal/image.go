package internal

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

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

	compressImage := resize.Resize(300, 0, img, resize.Bicubic)

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
