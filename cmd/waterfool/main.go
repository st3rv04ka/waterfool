package main

import (
	"flag"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	waterfoolImage "waterfool/v2/src/image"
)

func main() {
	imagePath := flag.String("image", "image.png", "image to convert")
	flag.Parse()

	img, err := waterfoolImage.ReadImage(*imagePath)
	if err != nil {
		log.Panic(err)
	}

	NFFT := img.Bounds().Max.X
	repetitions := waterfoolImage.Repetitions(NFFT)
	log.Println(repetitions)
	brightnessImg := waterfoolImage.ConvertImageToBrightnessArray(img, repetitions)
	brightnessImgReversed := waterfoolImage.FlipVertically(brightnessImg)
	waterfoolImage.MatrixToSignals(brightnessImgReversed, "output.iq8s", NFFT)
}
