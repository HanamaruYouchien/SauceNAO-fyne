package screen

import (
	"fyne.io/fyne/v2/container"
	"github.com/HanamaruYouchien/SauceNAO-fyne/pkg/saucenao"
	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
)

func ResultScreen(data *[]saucenao.Result) *container.Scroll {
	return component.ResultList(data)
}
