package app

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/qqMelon/silver-octo-guacamole/script"
)

func Window() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Silver Octo Guacamole")

	description := widget.NewLabel("Update, install, manage your World Of Warcraft addons with SOG !")

	str := binding.NewString()

	// Buttons
	elvuiBtn := widget.NewButton("Update/Install ElvUI", func() {
		err := str.Set("Update/Install ElvUI processing ...")
		if err != nil {
			return
		}
		err, resultVal := script.Elvui()
		err = str.Set(resultVal)
		if err != nil {
			return
		}
	})
	speedUpdateBtn := widget.NewButton("Speed Update", func() {
		err := str.Set("Start listing all addons ...")
		if err != nil {
			return
		}
		script.Update()
		str.Set("Games retrieved")
	})
	manageBtn := widget.NewButton("Manage Addons", func() {
		// WIP
		err := str.Set("Not available atm, on working ...")
		if err != nil {
			return
		}
	})

	// Dynamic data
	dynamicStr := widget.NewLabelWithData(str)

	// Container
	btnContainerCentered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), elvuiBtn, speedUpdateBtn, manageBtn, layout.NewSpacer())
	descriptionCentered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), description, layout.NewSpacer())
	dynamicStrCentered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), dynamicStr, layout.NewSpacer())

	mainContent := container.New(layout.NewVBoxLayout(), descriptionCentered, dynamicStrCentered, btnContainerCentered)

	myWindow.SetContent(mainContent)
	myWindow.ShowAndRun()
}
