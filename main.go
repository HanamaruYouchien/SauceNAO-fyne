package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	apikey := "myKey"
	w.SetContent(component.ApikeyEntry(apikey, func(newApikey string) {
		apikey = newApikey
		fmt.Println(apikey)
	}))

	image := &[]byte{}
	w2 := a.NewWindow("Select Image")
	w2.SetContent(component.ImageFileSelector(&w2, func(img *[]byte) {
		image = img
		fmt.Println(strconv.Itoa(len(*image)))
	}))

	w3 := a.NewWindow("Image Url")
	w3.SetContent(component.ImageUrlEntry())

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
	w4.ShowAndRun()
	_ = imageMethod
}
