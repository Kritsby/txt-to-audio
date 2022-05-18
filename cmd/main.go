package main

import (
	"fmt"
	"io/ioutil"
	"log"

	htgotts "github.com/hegedustibor/htgo-tts"
)

func main() {
	path := pathSave()

	text := readFile()

	speech := newSpeech(path)

	newAudioFile(speech, text)
}

func pathSave() (path string) {

	fmt.Print("Path for file save: ")
	fmt.Scan(&path)

	return path
}

func readFile() []byte {
	var fileName string

	fmt.Print("Path with text file with filename: ")
	fmt.Scan(&fileName)

	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("No such file in directory")
	}

	return text
}

func newSpeech(path string) htgotts.Speech {

	speech := htgotts.Speech{Folder: path, Language: "ru"}

	return speech
}

func newAudioFile(speech htgotts.Speech, text []byte) {

	speech.CreateSpeechFile(string(text), "audio")

}
