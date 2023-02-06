package app

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Window() {
	a := app.New()
	w := a.NewWindow("Silver Octo Guacamol")

	desc := widget.NewLabel("Update, Install, manage your addons with SOG")

	elvuiBtn := Button("Update/Install ElvUI", "updateElvui")
	updateBtn := Button("Speed update", "speedUpdate")
	manageBtn := Button("Manage addons", "manageAddon")

	btnContainer := container.New(layout.NewHBoxLayout(), elvuiBtn, updateBtn, manageBtn)
	desContainer := container.New(layout.NewHBoxLayout(), desc)

	//mainContent := container.New(layout.NewVBoxLayout(), desc, content, ActionStatus())

	centeredBtnContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnContainer, layout.NewSpacer())
	centeredDescContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), desContainer, layout.NewSpacer())

	w.SetContent(container.New(layout.NewVBoxLayout(), centeredDescContainer, centeredBtnContainer))

	w.ShowAndRun()
}
