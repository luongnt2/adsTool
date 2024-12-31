package builder

import "fyne.io/fyne/v2"

type Container interface {
	Name() string
	Size() fyne.Size
	Builder(parent fyne.Window) *fyne.Container
}
