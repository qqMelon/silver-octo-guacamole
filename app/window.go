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

	content := container.New(layout.NewHBoxLayout(), elvuiBtn, updateBtn, manageBtn)
	mainContent := container.New(layout.NewVBoxLayout(), desc, content, ActionStatus())

	w.SetContent(container.New(layout.NewCenterLayout(), mainContent))

	w.ShowAndRun()
}
