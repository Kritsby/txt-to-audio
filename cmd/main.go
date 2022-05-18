package main

import (
	"io/ioutil"
	"log"

	htgotts "github.com/hegedustibor/htgo-tts"
)

func main() {
	speech := htgotts.Speech{Folder: "/home/daniil/golang/projects/txt-to-audio/text-files", Language: "ru"}
	text, err := ioutil.ReadFile("/home/daniil/golang/projects/txt-to-audio/text-files/text.txt")
	if err != nil {
		log.Fatal("No such file in directory")
	}
	speech.Speak(string(text))
}
