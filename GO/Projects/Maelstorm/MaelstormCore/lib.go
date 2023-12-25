package maelstormcore

import (
	"fmt"
	"log"
)

func LogError(err error) {
	//todo: beautify
	if err != nil {
		log.Fatal(getRedBlinkString(err.Error()))
	}
}

func getRedBlinkString(message string) string {
	return getANSICode("\033[5;31m", message)
}

func getANSICode(code, message string) string {
	reset := "\033[0m"
	return fmt.Sprintf("%s%s%s", code, message, reset)
}
