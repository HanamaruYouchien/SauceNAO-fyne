package component

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

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
