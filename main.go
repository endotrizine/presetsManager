package main

import (
	"log"
	"presetsManager/tui"
)

func main() {
	err := tui.Start()
	if err != nil {
		log.Fatal(err)
	}
}
