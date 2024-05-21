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
	"github.com/HanamaruYouchien/SauceNAO-fyne/pkg/saucenao"
)

func ResultListByList(data *[]saucenao.Result) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*data)
		},
		func() fyne.CanvasObject {
			return widget.NewCard("", "", ResultItem(&(*data)[0]))
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Card).SetContent(ResultItem(&(*data)[lii]))
		},
	)
	return list
}

func ResultItem(data *saucenao.Result) *fyne.Container {
	var imgThumbnail *canvas.Image
	if rscThumbnail, err := fyne.LoadResourceFromURLString(data.Header.Thumbnail); err != nil {
		imgThumbnail = canvas.NewImageFromResource(theme.BrokenImageIcon())
	} else {
		imgThumbnail = canvas.NewImageFromResource(rscThumbnail)
	}
	imgThumbnail.FillMode = canvas.ImageFillContain
	imgThumbnail.SetMinSize(fyne.NewSize(100, 100))

	lbTitle := widget.NewLabel(data.GetTitle())
	lbDescription := widget.NewLabel(data.GetAuthor())

	var urlStr = data.GetUrls()[0]

	urlTarget, _ := url.Parse(urlStr)
	lnkTarget := widget.NewHyperlink(urlStr, urlTarget)
	ctnDetail := container.New(layout.NewVBoxLayout(), lbTitle, lbDescription, lnkTarget)

	lbSimilarity := widget.NewLabel(strconv.FormatFloat(data.Header.Similarity, 'f', 2, 64) + "%")
	sp := layout.NewSpacer()
	return container.New(layout.NewHBoxLayout(), imgThumbnail, ctnDetail, sp, lbSimilarity)
}

func ResultList(data *[]saucenao.Result) *container.Scroll {
	length := len(*data)
	var ctn *fyne.Container
	if length == 0 {
		ctn = container.New(layout.NewVBoxLayout(), widget.NewLabel("Not Found"))
	} else {
		ctn = container.New(layout.NewVBoxLayout(), ResultItem(&(*data)[0]))
	}
	for i := 1; i < length; i++ {
		ctn.Add(widget.NewSeparator())
		ctn.Add(ResultItem(&(*data)[i]))
	}

	scroll := container.NewVScroll(ctn)
	return scroll
}
