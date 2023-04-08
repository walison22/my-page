package pages

import "fyne.io/fyne/v2"

func setPage(win *fyne.Window, page string) {

	content := (*win).Content().(*fyne.Container)

	switch page {
	case "start":
		content.Objects[1] = Start(win)
	case "dashboard":
		content.Objects[1] = Dashboard(win)
	case "aboutme":
		content.Objects[1] = Aboutme(win)
	case "myskills":
		content.Objects[1] = Myskills(win)

	}
}
