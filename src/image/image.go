package image

import (
	"fmt"
	"image"
	"math"
	"math/cmplx"
	"math/rand"
	"os"

	"github.com/mjibson/go-dsp/fft"
)

const (
	freq     = 1000000
	lineTime = 0.01
)

func Repetitions(NFFT int) int {
	return int(float64(freq) * lineTime / float64(NFFT))
}

func prepareSpectrum(imageRow []float64, NFFT int) []complex128 {
	spectrum := make([]complex128, NFFT)

	for i, amp := range imageRow {
		phase := (float64(NFFT) + rand.Float64()*(float64(NFFT))) * 2 * math.Pi
		spectrum[i] = cmplx.Rect(amp, phase)
	}

	return spectrum
}

func FlipVertically(matrix [][]float64) [][]float64 {
	height := len(matrix)
	flippedMatrix := make([][]float64, height)

	for i, row := range matrix {
		flippedMatrix[height-i-1] = make([]float64, len(row))
		copy(flippedMatrix[height-i-1], matrix[i])
	}

	return flippedMatrix
}

func ReadImage(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ConvertImageToBrightnessArray(img image.Image, repetition int) [][]float64 {
	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy() * repetition
	brightnessMatrix := make([][]float64, imageHeight)

	for y := 0; y < imageHeight; y++ {
		brightnessMatrix[y] = make([]float64, imageWidth)
		for x := 0; x < imageWidth; x++ {
			originalY := (y / repetition) + img.Bounds().Min.Y
			r, g, b, _ := img.At(x, originalY).RGBA()
			brightness := (0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			brightnessMatrix[y][x] = brightness
		}
	}

	return brightnessMatrix
}

func MatrixToSignals(brightnessMatrix [][]float64, outputPath string, NFFT int) {
	file, err := os.Create(outputPath)
	if err != nil {
		return
	}
	defer file.Close()

	for index, imageRow := range brightnessMatrix {
		spectrum := prepareSpectrum(imageRow, NFFT)
		timeSignal := fft.IFFT(spectrum)
		if index == 0 {
			for _, i := range timeSignal {
				fmt.Print(real(i), ",")
			}
		}
		for _, val := range timeSignal {
			i := real(val)
			q := imag(val)
			file.Write([]byte{uint8(i), uint8(q)})
		}
	}
}
