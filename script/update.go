package script

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var baseUrl = "https://api.curseforge.com"

func getGame() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load env file ...")
	}

	headers := map[string][]string{
		"Accept":    []string{"application/json"},
		"Host":      {"api.curseforge.com"},
		"x-api-key": []string{os.Getenv("API_KEY")},
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
