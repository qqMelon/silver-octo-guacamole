package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
		if strings.Contains(e.Text, "13") {
			rv = e.Text
		}
	})

	c.Visit(baseUrl + "welcome.php")

	return rv
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please specify your ElvUI version you need")
		return
	}

	remoteVersion := checkRemoteVersion()
	fmt.Println(remoteVersion)
	version := os.Args[1]
	downloadPackage(version)
}
