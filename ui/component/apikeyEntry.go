package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ApikeyEntry(apikey string, onSave func(newApikey string)) *fyne.Container {
	entry := widget.NewPasswordEntry()
	entry.SetPlaceHolder("Input your API key")
	entry.Text = apikey

	btnSave := widget.NewButton("Save", func() {
		onSave(entry.Text)
	})
	btnContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnSave)
	return container.New(layout.NewVBoxLayout(), entry, btnContainer)
}
