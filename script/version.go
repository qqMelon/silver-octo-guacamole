package script

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func CheckRemoteVersion() string {
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

func CheckLocalVersion(version string) bool {
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
