package component

import (
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ResultItem(thumbnail string, title string, description string, link string, similarity float64) *fyne.Container {
	var imgThumbnail *canvas.Image
	if rscThumbnail, err := fyne.LoadResourceFromURLString(thumbnail); err != nil {
		imgThumbnail = canvas.NewImageFromResource(rscThumbnail)
	} else {
		imgThumbnail = canvas.NewImageFromResource(theme.BrokenImageIcon())
	}
	imgThumbnail.FillMode = canvas.ImageFillContain
	imgThumbnail.SetMinSize(fyne.NewSize(100, 100))

	lbTitle := widget.NewLabel(title)
	lbDescription := widget.NewLabel(description)
	urlTarget, _ := url.Parse(link)
	lnkTarget := widget.NewHyperlink(link, urlTarget)
	ctnDetail := container.New(layout.NewVBoxLayout(), lbTitle, lbDescription, lnkTarget)

	lbSimilarity := widget.NewLabel(strconv.FormatFloat(similarity, 'f', 2, 64) + "%")
	return container.New(layout.NewHBoxLayout(), imgThumbnail, ctnDetail, layout.NewSpacer(), lbSimilarity)
}
