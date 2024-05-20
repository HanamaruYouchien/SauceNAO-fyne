package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/HanamaruYouchien/SauceNAO-fyne/pkg/saucenao"
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

func ShowResultScreen(a *fyne.App, data *[]saucenao.Result) {
	resultScreen := (*a).NewWindow("Result")
	resultScreen.SetContent(screen.ResultScreen(data))
	resultScreen.Show()
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
			resp := &saucenao.Response{}
			var err error
			switch method {
			case component.METHOD_FILE:
				msg += "File\nFile size:" + strconv.Itoa(len(*img))
				resp, err = saucenao.SearchByFile(apikey, *img)
			case component.METHOD_URL:
				msg += "URL\nURL: " + url
				resp, err = saucenao.SearchByURL(apikey, url)
			}
			fmt.Println(msg)
			if err != nil {
				fmt.Println(err)
				return
			}

			ShowResultScreen(&a, &resp.Results)
		},
		func() {
			ShowSettingsScreen(&a, apikey, func(newApikey string) {
				apikey = newApikey
				pref.SetString(PREF_FIELD_APIKEY, apikey)
			})
		},
	))

	mainWindow.Show()
	a.Run()
}
