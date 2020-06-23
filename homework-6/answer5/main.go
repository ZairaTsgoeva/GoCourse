package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"os"
)

func createImage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "image/png")

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{}

	canvasSize := 400
	img := image.NewRGBA(image.Rect(0, 0, canvasSize, canvasSize))
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	cellSize := 50
	cellColor := black
	for i := 0; i <= canvasSize; i += cellSize {
		for j := 0; j <= canvasSize; j += cellSize {
			rect := image.Rect(i, j, i + cellSize, j + cellSize)
			draw.Draw(img, rect, &image.Uniform{cellColor}, image.Point{}, draw.Src)
			// может есть лучше решение для чередования цвета
			switch cellColor {
			case white:
				cellColor = black
			case black:
				cellColor = white
			}
		}
	}

	png.Encode(res, img)
}

func main() {
	http.HandleFunc("/create-image", createImage)
	http.ListenAndServe(":"+os.Args[1], nil)
}
