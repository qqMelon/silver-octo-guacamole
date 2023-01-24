package script

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadPackage(version string) {
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
