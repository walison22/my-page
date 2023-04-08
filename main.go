package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/me/content"
	"github.com/me/pages"
)

var windowDi = fyne.NewSize(960, 540)

func main() {
	a := app.New()
	win := a.NewWindow("My Profile")
	win.Resize(windowDi)
	win.SetPadded(false)
	win.CenterOnScreen()
	win.SetFixedSize(true)

	bg := content.Background()
	home := pages.Start(&win)

	maincontainer := container.NewMax(bg, home)

	win.SetContent(maincontainer)

	win.ShowAndRun()

}
