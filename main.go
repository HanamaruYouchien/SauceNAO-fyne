package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/screen"
)

const (
	PREF_FIELD_APIKEY = "apikey"
)

func ShowSettingsScreen(a *fyne.App, apikey string, onSave func(newApikey string)) {
	settingsScreen := (*a).NewWindow("Preferences")
	settingsScreen.SetContent(screen.SettingsScreen(apikey, onSave))
	settingsScreen.Show()
}

func main() {
	a := app.New()
	pref := a.Preferences()
	apikey := pref.StringWithFallback(PREF_FIELD_APIKEY, "")

	mainWindow := a.NewWindow("SauceNAO")

	mainWindow.SetMaster()
	mainWindow.SetContent(screen.MainScreen(
		&mainWindow,
		func(method bool, img *[]byte, url string) {
			msg := "Method: "
			switch method {
			case component.METHOD_FILE:
				msg += "File\nFile size:" + strconv.Itoa(len(*img))
			case component.METHOD_URL:
				msg += "URL\nURL: " + url
			}
			dialog.NewInformation("Query", msg, mainWindow).Show()
		},
		func() {
			ShowSettingsScreen(&a, apikey, func(newApikey string) {
				apikey = newApikey
				pref.SetString(PREF_FIELD_APIKEY, apikey)
				fmt.Println(apikey)
			})
		},
	))

	mainWindow.Show()
	a.Run()
}
