package main

import (
	"encoding/json"
	"fmt"
)

var (
	Init = "init"
)

type InputBody struct {
	Type    string   `json:"type"`
	NodeID  string   `json:"node_id"`
	NodeIds []string `json:"node_ids"`
	MsgID   int      `json:"msg_id"`
}

type InputMsg struct {
	ID   int       `json:"id"`
	Src  string    `json:"src"`
	Dest string    `json:"dest"`
	Body InputBody `json:"body"`
}

type OutputBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}

type OutputMsg struct {
	Dest string     `json:"dest"`
	Src  string     `json:"src"`
	Body OutputBody `json:"body"`
}

func logRecieved(msg InputMsg) {
	_json, err := json.Marshal(msg)
	LogError(err)
	GlobalOutPutHandler <- fmt.Sprintf("Received %s", _json)
}

func generateOutput(msg InputMsg) OutputMsg {
	return OutputMsg{
		Dest: msg.Src,
		Src:  msg.Dest,
	}
}

func Dispatch(data string) {
	msg := InputMsg{}
	LogError(json.Unmarshal([]byte(data), &msg))
	logRecieved(msg)
	switch msg.Body.Type {
	case Init:
		Nodes.AddNode(msg.Body, generateOutput(msg))
	}
}
