package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
)

const (
	PREF_FIELD_APIKEY = "apikey"
)

func main() {
	a := app.New()
	pref := a.Preferences()

	image := &[]byte{}
	w4 := a.NewWindow("ImagePicker")
	imageMethod := component.METHOD_FILE
	w4.SetContent(component.ImagePicker(&w4, func(method bool) {
		imageMethod = method
	}, func(img *[]byte) {
		image = img
		fmt.Println(strconv.Itoa(len(*image)))
	}, func() {
		msg := "Method: "
		switch imageMethod {
		case component.METHOD_FILE:
			msg += "File\nFile size:" + strconv.Itoa(len(*image))
		case component.METHOD_URL:
			msg += "URL"
		}
		dialog.NewInformation("Query", msg, w4).Show()
	}))

	apikey := pref.StringWithFallback(PREF_FIELD_APIKEY, "")
	w := a.NewWindow("Hello World")
	w.SetContent(component.ApikeyEntry(apikey, func(newApikey string) {
		apikey = newApikey
		pref.SetString(PREF_FIELD_APIKEY, apikey)
		fmt.Println(apikey)
	}))

	w2 := a.NewWindow("Select Image")
	w2.SetContent(component.ImageFileSelector(&w2, func(img *[]byte) {
		image = img
		fmt.Println(strconv.Itoa(len(*image)))
	}))

	w3 := a.NewWindow("Image Url")
	w3.SetContent(component.ImageUrlEntry())

	w.ShowAndRun()
}
