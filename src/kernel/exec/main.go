package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	// err := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "New Contents Avairable!!" sound name "Hero"`, createText())).Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	command := `
	tell application "Safari"
		open location "%s"
		activate
	end tell
	`
	for _, link := range []string{"https://yurucamp.jp/", "https://yurucamp.jp/second/"} {
		err := exec.Command("osascript", "-e", fmt.Sprintf(command, link)).Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	// err = exec.Command("afplay", "/System/Library/Sounds/Hero.aiff").Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func createText() string {
	sample0Text := strings.Join([]string{"sample0", "http"}, "\n\t")
	sample1Text := strings.Join([]string{"sample1", "http"}, "\n\t")
	sample2Text := strings.Join([]string{"sample2", "http"}, "\n\t")
	return strings.Join([]string{sample0Text, sample1Text, sample2Text}, "\n")
}
