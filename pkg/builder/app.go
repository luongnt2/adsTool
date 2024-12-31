package builder

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var myApp fyne.App

func init() {
	myApp = app.New()
}

func NewWindow(container Container) fyne.Window {
	window := myApp.NewWindow(container.Name())
	window.SetContent(container.Builder(window))
	window.Resize(container.Size())

	return window
}
