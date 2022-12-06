package main

import (
	"log"
	"os"
)

type writerLog interface {
	log(data string)
}

type WriterLog struct {
	writerLog
	filename string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (w *WriterLog) log(data string) {
	err := os.WriteFile(w.filename, []byte(data), 0644)
	checkError(err)
}

type errorWriterLog interface {
	log(data string)
}

type ErrorWriterLog struct {
	errorWriterLog
	logger WriterLog
}

func (e *ErrorWriterLog) log(data string) {
	e.logger.log("**********************\n Failed \n\n\n\n" + data + "\n\n\n\n **********************")
}

func main() {
	writer := WriterLog{
		filename: "TesterDecorator.txt",
	}
	errorWriter := ErrorWriterLog{
		logger: writer,
	}

	errorWriter.log("Just for testing Decorator pattern")

}
