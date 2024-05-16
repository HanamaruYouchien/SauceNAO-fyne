package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/HanamaruYouchien/SauceNAO-fyne/ui/component"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(component.ApikeyEntry(func() {}))
	w.ShowAndRun()
}
