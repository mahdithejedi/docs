package helpers

import (
	"fmt"
	"io"
	"os"
)

func PureCheckErr(err error) {
	CheckErr("", err)
}

func CheckErr(extra string, err error) {
	if err != nil {
		RaiseStdErr(extra + " : " + err.Error())
	}
}

func RaiseStdErr(error string) {
	RaiseException(error, os.Stderr)
}

func RaiseException(error string, w io.Writer) {
	fmt.Fprintln(w, error)
	os.Exit(0)
}
