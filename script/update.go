package script

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var baseUrl = "https://api.curseforge.com"

func getGame() {
	headers := map[string][]string{
		"Accept":    {"application/json"},
		"x-api-key": {os.Getenv("API_KEY")},
	}

	data := bytes.NewBuffer([]byte{})
	req, _ := http.NewRequest("GET", baseUrl+"/v1/games", data)
	req.Header = headers

	client := &http.Client{}
	resp, _ := client.Do(req)

	fmt.Println(resp)
}

func Update() {
	getGame()
}
