package main

import (
	"fmt"

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
	w.ShowAndRun()
}
