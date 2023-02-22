package script

import (
	"archive/zip"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func checkRemoteVersion() string {
	var rv = ""
	var baseUrl = "https://www.tukui.org/"

	c := colly.NewCollector(
		colly.AllowedDomains("www.tukui.org"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Check url: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status: ", r.StatusCode)
	})

	c.OnHTML(".hidden-xs", func(e *colly.HTMLElement) {
		// Find func to catch ElvUI only and update this func / maybe update major version if needed
		if strings.Contains(e.Text, "13") {
			rv = e.Text
		}
	})

	err := c.Visit(baseUrl + "welcome.php")
	if err != nil {
		return ""
	}

	return rv
}

func checkLocalVersion(version string) bool {
	_, err := os.Stat("elvui-" + version + ".zip")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err)
			return false
		}
		return false
	}
	return true
}

func downloadPackage(version string) {
	var baseUrl string = "https://www.tukui.org/"
	url := fmt.Sprintf(baseUrl+"downloads/elvui-%s", version+".zip")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in getting url: ", err)
		return
	}

	defer resp.Body.Close()

	filename := fmt.Sprintf("elvui-%s.zip", version)
	out, _ := os.Create(filename)

	defer out.Close()

	io.Copy(out, resp.Body)
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Elvui() (error, string) {
	remoteVersion := checkRemoteVersion()
	fmt.Println("Remote version: ", remoteVersion)

	if checkLocalVersion(remoteVersion) {
		return nil, fmt.Sprintf("You already have the latest version of ElvUI: %s", remoteVersion)
	}

	downloadPackage(remoteVersion)
	unzip("elvui-"+remoteVersion+".zip", "Addons/")

	return nil, fmt.Sprintf("Successfully download and install ElvUI: %s", remoteVersion)
}
