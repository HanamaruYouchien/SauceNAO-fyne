package screen

import (
	"fyne.io/fyne/v2"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
)

func SettingsScreen(apikey string, onSave func(newApikey string)) *fyne.Container {
	return component.ApikeyEntry(apikey, onSave)
}
