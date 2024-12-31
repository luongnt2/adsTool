package containers

import (
	"adsTool/pkg/builder"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Home struct {
}

func (c *Home) Name() string {
	return "Home"
}

func (c *Home) Size() fyne.Size {
	return fyne.NewSize(400, 400)
}

func (c *Home) Builder(fyne.Window) *fyne.Container {
	table := widget.NewTable(
		func() (int, int) { return 10, 2 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			//label := cell.(*widget.Label)
			//if id.Col == 0 {
			//	label.SetText(getName(id.Row))
			//} else {
			//	label.SetText(getEmail(id.Row))
			//}
		},
	)

	addWindow := builder.NewWindow(&Add{})
	addButton := widget.NewButton("Thêm hành đông chuyển đổi", func() {
		addWindow.Show()
	})

	return container.NewBorder(nil, addButton, nil, nil, table)
}
