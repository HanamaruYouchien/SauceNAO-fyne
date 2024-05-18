package component

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ToolbarMain(openSettings func()) *widget.Toolbar {
	toolbar := widget.NewToolbar(widget.NewToolbarSpacer(), widget.NewToolbarAction(theme.SettingsIcon(), openSettings))

	return toolbar
}
