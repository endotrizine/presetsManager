package main

import (
	"log"
	"presetsManager/tui"
)

func main() {
	err := tui.StartThemes()
	if err != nil {
		log.Fatal(err)
	}
}
