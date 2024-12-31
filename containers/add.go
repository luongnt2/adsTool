package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shenzhencenter/google-ads-pb/enums"
)

var types []string

func init() {
	for _, t := range enums.ConversionActionTypeEnum_ConversionActionType_name {
		types = append(types, t)
	}
}

type Add struct {
}

func (c *Add) Name() string {
	return "Add new"
}

func (c *Add) Size() fyne.Size {
	return fyne.NewSize(400, 200)
}

func (c *Add) Builder(parent fyne.Window) *fyne.Container {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter name")

	typeEntry := widget.NewSelect(types, func(s string) {

	})

	submitButton := widget.NewButton("Submit", func() {
		//name := nameEntry.Text
		email := typeEntry.Selected
		print(email)
		//if name != "" && email != "" {
		//	_, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
		//	if err != nil {
		//		dialog.ShowError(err, parent)
		//	} else {
		//		parent.Close()
		//	}
		//}
	})

	return container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("Name", nameEntry),
			widget.NewFormItem("Type", typeEntry),
		),
		submitButton,
	)
}
