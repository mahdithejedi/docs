package utils

import (
	"fmt"
	"log"
)

func CheckError(err error) {
	CheckErrorWithMessage(err, "")
}

func CheckErrorWithMessage(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprintf(
			"Error %s : %s", msg, err.Error(),
		))
	}
}
