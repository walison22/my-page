package content

import (
	_ "embed"

	"fyne.io/fyne/v2/canvas"
)

func Background() *canvas.Image {
	image := canvas.NewImageFromFile("C:/Users/Emmanuel Wali Okocha/Desktop/bg/Comp 1/mypage_00000.png")
	image.FillMode = canvas.ImageFillStretch
	image.ScaleMode = canvas.ImageScaleFastest

	return image
}
