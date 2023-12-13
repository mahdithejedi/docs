package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var Terminal = make(chan string)
var SignalChan = make(chan os.Signal, 1)

func initTerminalHandler() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Terminal <- scanner.Text()
	}
}

func initSignalHandler() {
	signal.Notify(SignalChan, syscall.SIGINT)
}

func main() {
	go initTerminalHandler()
	initSignalHandler()
	go Nodes.InitOutputHandler()
	defer func() {
		close(Terminal)
		GlobalOutPutHandler <- "0"
	}()
	for {
		select {
		case data := <-Terminal:
			if json.Valid([]byte(data)) {
				Dispatch(data)
			} else {
				GlobalOutPutHandler <- fmt.Sprintf("Received %s !!!!!!!!!!!!!@@@@@@@@@@@@@############", data)
			}
		case <-SignalChan:
			return
		}
	}

}
