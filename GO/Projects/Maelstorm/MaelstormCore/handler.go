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

type HandlerFunction func(nodeID string, message InputBody) OutputBody

//type HandlerChannelFunction func(nodeID string, message chan InputBody) chan OutputBody

//func(string, InputBody) OutputBody type

type MsgInputType string
type MsgOutputType string

var (
	Init      MsgInputType = "init"
	Echo      MsgInputType = "echo"
	Topology  MsgInputType = "topology"
	Broadcast MsgInputType = "broadcast"
)

var (
	InitOk      MsgOutputType = "init_ok"
	EchoOk      MsgOutputType = "echo_ok"
	TopologyOK  MsgOutputType = "topology_ok"
	BroadcastOK MsgOutputType = "broadcast_ok"
)

type IHandler interface {
	AddHandler(MsgInputType, HandlerFunction)
	//AddChannelHandler(MsgInputType, HandlerChannelFunction)
}

type Handler struct {
	IHandler
	handlers map[MsgInputType]HandlerFunction
	Input    io.Reader
	Output   io.Writer
	AutoInit bool
}

func (h *Handler) Init() {
	if h.AutoInit {
		h.AddHandler(Init, func(nodeID string, message InputBody) (output OutputBody) {
			if Nodes.AddNode(nodeID) == true {
				GlobalLogChannel <- fmt.Sprintf("Node %s initiated while all nodes are %#v", message.NodeID, Nodes.nodes)
			}
			output = OutputBody{
				Type:      string(InitOk),
				InReplyTo: message.MsgID,
			}
			return
		})
	}
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
	m.Output.BodyStruct = h.handlers[m.GetMessageType()](m.Input.Dest, m.Input.BodyStruct)
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
	InitNodeHandler()
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

func (h *Handler) AddHandler(name MsgInputType, function HandlerFunction) {
	h.handlers[name] = function
}

func InitHandler(autoInit bool) *Handler {
	_handler := &Handler{
		handlers: make(map[MsgInputType]HandlerFunction),
		Input:    os.Stdin,
		Output:   os.Stdout,
		AutoInit: autoInit,
	}
	_handler.Init()
	return _handler
}
