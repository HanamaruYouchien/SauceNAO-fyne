package component

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

const (
	METHOD_FILE = false
	METHOD_URL  = true
)

func ImagePicker(win *fyne.Window, setMethod func(bool), onSelectImage func(image *[]byte), onSearch func()) *fyne.Container {
	const (
		OPTION_FILE = "File"
		OPTION_URL  = "URL"
	)
	title := widget.NewLabel("Search by")

	selector := ImageFileSelector(win, onSelectImage)

	entry := ImageUrlEntry()

	radio := widget.NewRadioGroup([]string{OPTION_FILE, OPTION_URL}, func(val string) {
		state := true
		switch val {
		case OPTION_FILE:
			state = METHOD_FILE
		case OPTION_URL:
			state = METHOD_URL
		}
		setMethod(state)
		selector.Hidden = state
		entry.Hidden = !state
	})
	radio.Horizontal = true
	radio.SetSelected(OPTION_FILE)
	switchMethodContainer := container.New(layout.NewHBoxLayout(), title, layout.NewSpacer(), radio)

	btnSearch := widget.NewButton("Search", onSearch)

	return container.New(layout.NewVBoxLayout(), switchMethodContainer, selector, entry, btnSearch)
}

func ImageUrlEntry() *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Image URL to search")
	// entry.OnSubmitted
	return entry
}

func ImageFileSelector(win *fyne.Window, onSelectImage func(image *[]byte)) *widget.Button {
	btnOpenImage := widget.NewButton("Select Image", func() {
		imageDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, *win)
				return
			}
			if reader == nil {
				dialog.ShowInformation("Cancelled", "Cancelled by user", *win)
				return
			}
			defer reader.Close()

			img, err := io.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, *win)
				return
			}
			onSelectImage(&img)
		}, *win)
		imageDialog.SetFilter(storage.NewMimeTypeFileFilter([]string{"image/*"}))
		imageDialog.Show()
	})
	return btnOpenImage
}
