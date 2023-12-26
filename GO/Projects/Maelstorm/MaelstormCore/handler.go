package maelstormcore

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var MaxBufferSize = 120
var GlobalLogChannel = make(chan string, MaxBufferSize)
var GlobalWriterChannel = make(chan []byte, MaxBufferSize)
var Terminal = make(chan string)
var SignalChan = make(chan os.Signal, 1)

type HandlerFunction func(message InputBody) OutputBody

type IHandler interface {
	AddHandler(MsgType, HandlerFunction)
}

type MsgType string

var (
	InitOk MsgType = "init_ok"
	Init   MsgType = "init"
)

type Handler struct {
	IHandler
	handlers map[MsgType]HandlerFunction
	Input    io.Reader
	Output   io.Writer
}

func (h *Handler) initTerminalHandler() {
	scanner := bufio.NewScanner(h.Input)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(strings.TrimSpace(txt)) > 0 {
			Terminal <- txt
		}
	}
}

func (h *Handler) initSignalHandler() {
	signal.Notify(SignalChan, syscall.SIGINT)
}

func (h *Handler) initLogHandler() {
	for {
		msg := <-GlobalLogChannel
		if msg == "0" {
			break
		}
		log.Printf(msg)
	}
	defer close(GlobalLogChannel)
}

func (h *Handler) initWriter() {
	for {
		msg := <-GlobalWriterChannel
		if string(msg) == "0" {
			break
		}
		_, err := h.Output.Write(msg)
		LogError(err)
		_, err = h.Output.Write([]byte("\n"))
		LogError(err)
	}
	defer close(GlobalWriterChannel)
}

func (h *Handler) run() {
	for {
		select {
		case data := <-Terminal:
			if json.Valid([]byte(data)) {
				h.process(data)
			} else if data == "" {
				continue
			} else {
				LogError(
					errors.New(
						fmt.Sprintf("Packet %s can not be processes because it's not be converted in JSON",
							data),
					),
				)
			}
		case <-SignalChan:
			return
		}
	}
}

func (h *Handler) process(data string) {
	m := Dispatch(data)
	GlobalLogChannel <- fmt.Sprintf("Received %s", data)
	m.Output.Dest = m.Input.Src
	m.Output.Src = m.Input.Dest
	m.Output.BodyStruct = h.handlers[m.GetMessageType()](m.Input.BodyStruct)
	h.reply(m.Marshal())
}

func (h *Handler) reply(msg []byte) {
	GlobalWriterChannel <- msg
}

func (h *Handler) destruct() {
	close(Terminal)
	GlobalLogChannel <- "0"
	GlobalWriterChannel <- []byte("0")
}

func (h *Handler) init() {
	go h.initTerminalHandler()
	h.initSignalHandler()
	go h.initLogHandler()
	go h.initWriter()
}

func (h *Handler) Run() {
	h.init()
	h.run()
	defer h.destruct()
}

func (h *Handler) AddHandler(name MsgType, function HandlerFunction) {
	h.handlers[name] = function
}

func InitHandler() *Handler {
	return &Handler{
		handlers: make(map[MsgType]HandlerFunction),
		Input:    os.Stdin,
		Output:   os.Stdout,
	}
}
