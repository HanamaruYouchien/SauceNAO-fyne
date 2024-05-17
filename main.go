package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"

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
	w3.ShowAndRun()
}
