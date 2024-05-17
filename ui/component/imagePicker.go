package component

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

const (
	METHOD_FILE = false
	METHOD_URL  = true
)

func ImagePicker(win *fyne.Window, onSearch func(method bool, img *[]byte, url string)) *fyne.Container {
	const (
		OPTION_FILE = "File"
		OPTION_URL  = "URL"
	)

	title := widget.NewLabel("Search by")

	image := &[]byte{}
	selector := ImageFileSelector(win, func(img *[]byte) {
		image = img
	})

	etyUrl := binding.NewString()
	entry := ImageUrlEntry(&etyUrl)

	method := METHOD_FILE
	radio := widget.NewRadioGroup([]string{OPTION_FILE, OPTION_URL}, func(val string) {
		switch val {
		case OPTION_FILE:
			method = METHOD_FILE
		case OPTION_URL:
			method = METHOD_URL
		}
		selector.Hidden = method
		entry.Hidden = !method
	})
	radio.Horizontal = true
	radio.SetSelected(OPTION_FILE)
	switchMethodContainer := container.New(layout.NewHBoxLayout(), title, layout.NewSpacer(), radio)

	btnSearch := widget.NewButton("Search", func() {
		url, err := etyUrl.Get()
		if err != nil {
			return
		}
		onSearch(method, image, url)
	})

	return container.New(layout.NewVBoxLayout(), switchMethodContainer, selector, entry, btnSearch)
}

func ImageUrlEntry(data *binding.String) *widget.Entry {
	entry := widget.NewEntryWithData(*data)
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
