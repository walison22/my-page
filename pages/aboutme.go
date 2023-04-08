package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Aboutme(win *fyne.Window) fyne.CanvasObject {

	title2 := canvas.NewText("ABOUT ME", color.White)
	title2.Alignment = fyne.TextAlignCenter
	title2.TextStyle.Bold = true
	title2.TextSize = 28
	title2.Resize(fyne.NewSize(300, 50))
	title2.Move(fyne.NewPos(320, 150))

	line1 := canvas.NewLine(color.RGBA{R: 255, G: 165, B: 0, A: 255})
	line1.Resize(fyne.NewSize(800, 100))
	line1.Position1 = fyne.NewPos(0, 125)
	line1.Position2 = fyne.NewPos(1750, 125)
	line1.StrokeWidth = 3

	Text4 := canvas.NewText("DEDICATED, HARDWORKING AND TALENTED GRADUATE OF", color.White)
	Text4.Alignment = fyne.TextAlignCenter
	Text4.TextStyle = fyne.TextStyle{Bold: true}
	Text4.TextSize = 15

	Text5 := canvas.NewText("CHEMICAL ENGINEERING WITH IN DEPTH KNOWLEDGE IN", color.White)
	Text5.Alignment = fyne.TextAlignCenter
	Text5.TextStyle = fyne.TextStyle{Bold: true}
	Text5.TextSize = 15

	Text6 := canvas.NewText("SOFTWARE ENGINEERING, VIDEO EDITING, LOGO ANIMATION", color.White)
	Text6.Alignment = fyne.TextAlignCenter
	Text6.TextStyle = fyne.TextStyle{Bold: true}
	Text6.TextSize = 15

	Text7 := canvas.NewText("AND GRAPHICS DESIGN. OVER THE YEARS I HAVE WORKED WITH", color.White)
	Text7.Alignment = fyne.TextAlignCenter
	Text7.TextStyle = fyne.TextStyle{Bold: true}
	Text7.TextSize = 15

	Text8 := canvas.NewText("MANY PROJECTS ACROSS THE GLOBE GIVING ME A WIDE RANGE OF EXPERIENCE.", color.White)
	Text8.Alignment = fyne.TextAlignCenter
	Text8.TextStyle = fyne.TextStyle{Bold: true}
	Text8.TextSize = 15

	backbtn := widget.NewButton("Back", func() {
		setPage(win, "dashboard")
	})
	backbtn.Move(fyne.NewPos(130, 400))
	backbtn.Importance = widget.WarningImportance
	backbtn.Resize(fyne.NewSize(250, 250))

	image1 := canvas.NewImageFromFile("C:/Users/Emmanuel Wali Okocha/Desktop/bg/Comp 1/photo_00000.png")
	image1.FillMode = canvas.ImageFillOriginal
	image1.Resize(fyne.NewSize(1500, 590))
	image1.Move(fyne.NewPos(-250, -170))

	n := container.NewHBox(backbtn)
	n.Move(fyne.NewPos(455, 400))

	t := container.NewVBox(Text4, Text5, Text6, Text7, Text8)
	t.Move(fyne.NewPos(170, 250))

	r := container.NewWithoutLayout(n, image1, t, title2, line1)

	return r

}
