package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Myskills(win *fyne.Window) fyne.CanvasObject {

	title9 := canvas.NewText("MY SKILLS", color.White)
	title9.Alignment = fyne.TextAlignCenter
	title9.TextStyle.Bold = true
	title9.TextSize = 28
	title9.Resize(fyne.NewSize(300, 50))
	title9.Move(fyne.NewPos(320, 150))

	line9 := canvas.NewLine(color.RGBA{R: 255, G: 165, B: 0, A: 255})
	line9.Resize(fyne.NewSize(800, 100))
	line9.Position1 = fyne.NewPos(0, 125)
	line9.Position2 = fyne.NewPos(1750, 125)
	line9.StrokeWidth = 3

	swebtn := widget.NewButton("SOFTWARE ENGINEER", func() {})
	swebtn.Importance = widget.WarningImportance
	swebtn.Resize(fyne.NewSize(250, 45))
	swebtn.Move(fyne.NewPos(200, 250))

	vebtn := widget.NewButton("VIDEO EDITOR", func() {})
	vebtn.Importance = widget.WarningImportance
	vebtn.Resize(fyne.NewSize(250, 45))
	vebtn.Move(fyne.NewPos(500, 250))

	labtn := widget.NewButton("GRAPHICS DESIGNER", func() {})
	labtn.Importance = widget.WarningImportance
	labtn.Resize(fyne.NewSize(250, 45))
	labtn.Move(fyne.NewPos(200, 350))

	cmbtn := widget.NewButton("LOGO ANIMATOR", func() {})
	cmbtn.Importance = widget.WarningImportance
	cmbtn.Resize(fyne.NewSize(250, 45))
	cmbtn.Move(fyne.NewPos(500, 350))

	bakbtn := widget.NewButton("Back", func() {
		setPage(win, "dashboard")
	})
	bakbtn.Move(fyne.NewPos(415, 410))
	bakbtn.Importance = widget.WarningImportance
	bakbtn.Resize(fyne.NewSize(120, 30))

	image9 := canvas.NewImageFromFile("C:/Users/Emmanuel Wali Okocha/Desktop/bg/Comp 1/photo_00000.png")
	image9.FillMode = canvas.ImageFillOriginal
	image9.Resize(fyne.NewSize(1500, 590))
	image9.Move(fyne.NewPos(-250, -170))

	p := container.NewWithoutLayout(image9, bakbtn, swebtn, vebtn, labtn, cmbtn, line9, title9)
	return p
}
