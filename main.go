package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/screen"
)

const (
	PREF_FIELD_APIKEY = "apikey"
)

func main() {
	a := app.New()
	pref := a.Preferences()

	w4 := a.NewWindow("ImagePicker")
	w4.SetContent(screen.MainScreen(&w4, func(method bool, img *[]byte, url string) {
		msg := "Method: "
		switch method {
		case component.METHOD_FILE:
			msg += "File\nFile size:" + strconv.Itoa(len(*img))
		case component.METHOD_URL:
			msg += "URL\nURL: " + url
		}
		dialog.NewInformation("Query", msg, w4).Show()
	}, func() { fmt.Println("settings") }))

	apikey := pref.StringWithFallback(PREF_FIELD_APIKEY, "")
	w := a.NewWindow("Hello World")
	w.SetContent(component.ApikeyEntry(apikey, func(newApikey string) {
		apikey = newApikey
		pref.SetString(PREF_FIELD_APIKEY, apikey)
		fmt.Println(apikey)
	}))

	w4.ShowAndRun()
}
