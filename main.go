package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/qqMelon/silver-octo-guacamole/script"
)

func UITool() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}

func main() {
	remoteVersion := script.CheckRemoteVersion()
	fmt.Println("Remote version: ", remoteVersion)

	UITool()

	if script.CheckLocalVersion(remoteVersion) {
		fmt.Println("You already have the latest version of ElvUI")
		time.Sleep(3 * time.Second)
		return
	}
	fmt.Println("Downloading ElvUI...")
	script.DownloadPackage(remoteVersion)
	fmt.Println("Download complete")
	fmt.Println("Extracting...")
	script.Unzip("elvui-"+remoteVersion+".zip", "Addons/")
	fmt.Println("Extracting done at : Addons/")
	time.Sleep(3 * time.Second)
}
