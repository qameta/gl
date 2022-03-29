package main

import (
	"log"
	"qameta.io/gl/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
