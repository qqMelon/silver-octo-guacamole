package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirs, err := ioutil.ReadDir("./../AddOns/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of Addons: ")
	for _, f := range dirs {
		fmt.Println(f.Name())
	}
}
