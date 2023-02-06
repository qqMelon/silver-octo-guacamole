package app

import (
	"fyne.io/fyne/v2/widget"
	"github.com/qqMelon/silver-octo-guacamole/script"
	"log"
)

func Button(title, action string) *widget.Button {
	content := widget.NewButton(title, func() {
		log.Println(action)
		if action == "updateElvui" {
			log.Println("Execute elvuiupdate function")
			script.Elvui()
		} else if action == "speedUpdate" {
			log.Println("Execute speedUpdate function")
		} else {
			log.Println("Switch to a new view ofr manage Addons")
		}
	})
	return content
}
