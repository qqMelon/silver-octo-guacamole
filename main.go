package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

var (
	baseUrl string = "https://www.tukui.org/"
)

func downloadPackage(version string) {
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

func checkRemoteVersion() string {
	var rv string = ""

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

	c.Visit(baseUrl + "welcome.php")

	return rv
}

func Unzip(src, dest string) error {
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

func main() {
	remoteVersion := checkRemoteVersion()
	fmt.Println(remoteVersion)

	if checkLocalVersion(remoteVersion) {
		fmt.Println("You already have the latest version of ElvUI")
		return
	}
	downloadPackage(remoteVersion)
	Unzip("elvui-"+remoteVersion+".zip", "Addons/")
}
