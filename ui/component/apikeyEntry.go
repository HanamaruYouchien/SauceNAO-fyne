package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ApikeyEntry(onSave func()) *fyne.Container {
	entry := widget.NewPasswordEntry()
	entry.SetPlaceHolder("Input your API key")

	btnSave := widget.NewButton("Save", onSave)
	btnContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnSave)
	return container.New(layout.NewVBoxLayout(), entry, btnContainer)
}
