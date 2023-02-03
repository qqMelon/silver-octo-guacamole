package main

import (
	"fmt"
	"time"

	"github.com/qqMelon/silver-octo-guacamole/script"
)

func main() {
	remoteVersion := script.CheckRemoteVersion()
	fmt.Println("Remote version: ", remoteVersion)

	if script.CheckLocalVersion(remoteVersion) {
		fmt.Println("You already have the latest version of ElvUI")
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
