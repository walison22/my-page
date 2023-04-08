package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Dashboard(win *fyne.Window) fyne.CanvasObject {

	title1 := canvas.NewText("Wali Emmanuel Okocha", color.White)
	title1.Alignment = fyne.TextAlignCenter
	title1.TextStyle.Bold = true
	title1.TextSize = 20
	title1.Resize(fyne.NewSize(300, 50))
	title1.Move(fyne.NewPos(350, 150))

	Text := canvas.NewText("Welcome To My Page, Here you will Find Information About Me,My Skills, My Works and how To reach me", color.White)
	Text.Alignment = fyne.TextAlignTrailing
	Text.TextStyle = fyne.TextStyle{Bold: true}
	Text.TextSize = 15
	Text.Resize(fyne.NewSize(100, 30))
	Text.Move(fyne.NewPos(750, 260))

	Text1 := canvas.NewText("This Website was created by me, i used a programming language called GOLANG, and a framework in it called FYNE ", color.White)
	Text1.Alignment = fyne.TextAlignTrailing
	Text1.TextStyle = fyne.TextStyle{Bold: true}
	Text1.TextSize = 15
	Text1.Resize(fyne.NewSize(100, 30))
	Text1.Move(fyne.NewPos(810, 300))

	line := canvas.NewLine(color.RGBA{R: 255, G: 165, B: 0, A: 255})
	line.Resize(fyne.NewSize(800, 100))
	line.Position1 = fyne.NewPos(0, 85)
	line.Position2 = fyne.NewPos(1750, 85)
	line.StrokeWidth = 3

	image := canvas.NewImageFromFile("C:/Users/Emmanuel Wali Okocha/Desktop/bg/Comp 1/photo_00000.png")
	image.FillMode = canvas.ImageFillOriginal
	image.Resize(fyne.NewSize(1750, 250))
	image.Move(fyne.NewPos(10, -50))

	aboutmebtn := widget.NewButton("About Me", func() {
		setPage(win, "aboutme")
	})
	aboutmebtn.Move(fyne.NewPos(400, 5))
	aboutmebtn.Importance = widget.WarningImportance

	myskillsbtn := widget.NewButton("My Skills", func() {
		setPage(win, "myskills")
	})
	myskillsbtn.Importance = widget.WarningImportance
	myworksbtn := widget.NewButton("My Works", func() {})
	myworksbtn.Importance = widget.WarningImportance

	c := container.NewHBox(container.NewCenter(aboutmebtn), container.NewCenter(myskillsbtn), container.NewCenter(myworksbtn))
	c.Resize(fyne.NewSize(100, 100))
	c.Move(fyne.NewPos(45, 0))
	d := container.NewWithoutLayout(Text, Text1, title1, line, image, c)

	return d

}
