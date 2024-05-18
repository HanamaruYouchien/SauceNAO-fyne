package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
)

func MainScreen(
	win *fyne.Window,
	onSearch func(method bool, img *[]byte, url string),
	openSettings func(),
) *fyne.Container {
	return container.New(layout.NewVBoxLayout(), component.ToolbarMain(openSettings), component.ImagePicker(win, onSearch))
}
