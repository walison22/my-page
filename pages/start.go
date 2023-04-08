package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var gap = layout.NewSpacer()

func Start(win *fyne.Window) *fyne.Container {

	title := canvas.NewText("Welcome to my Page", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 30
	title.Resize(fyne.NewSize(300, 50))
	title.Move(fyne.NewPos(300, 50))

	Proceedbtn := widget.NewButton("Proceed to HomePage", func() {
		setPage(win, "dashboard")

	})
	Proceedbtn.Importance = widget.WarningImportance

	vcontainer := container.NewVBox(title, Proceedbtn, gap)
	vcontainer.Resize(fyne.NewSize(400, 400))
	vcontainer.Move(fyne.NewPos(280, 200))

	return container.NewWithoutLayout(vcontainer)
}
