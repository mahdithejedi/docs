package maelstormcore

import (
	"encoding/json"
	"fmt"
)

type InputBody struct {
	Type    string   `json:"type"`
	NodeID  string   `json:"node_id,omitempty"`
	NodeIds []string `json:"node_ids,omitempty"`
	MsgID   int      `json:"msg_id"`
	// Instruction
	// Echo
	Echo string `json:"echo,omitempty"`
	// Topology
	Topology map[string][]string `json:"topology,omitempty"`
}

type InputMsg struct {
	ID         int             `json:"id"`
	Src        string          `json:"src"`
	Dest       string          `json:"dest"`
	Body       json.RawMessage `json:"body"`
	BodyStruct InputBody       `json:"-"`
}

type OutputBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
	// Instructions
	MsgID int    `json:"msg_id,omitempty"`
	Echo  string `json:"echo,omitempty"`
}

type OutputMsg struct {
	Dest       string          `json:"dest"`
	Src        string          `json:"src"`
	Body       json.RawMessage `json:"body"`
	BodyStruct OutputBody      `json:"-"`
}

type Message struct {
	Input  InputMsg
	Output OutputMsg
}

func (i *InputMsg) String() string {
	return fmt.Sprintf("%s", struct {
		Src  string
		Dest string
		Body json.RawMessage
	}{
		Src:  i.Src,
		Dest: i.Dest,
		Body: i.Body,
	})

}

func (m *Message) unpack(data string) {
	LogError(json.Unmarshal([]byte(data), &m.Input))
	inputBody := InputBody{}
	LogError(json.Unmarshal(m.Input.Body, &inputBody))
	m.Input.BodyStruct = inputBody
}

func (m *Message) GetMessageType() MsgInputType {
	return MsgInputType(m.Input.BodyStruct.Type)
}

func (m *Message) Marshal() (response []byte) {
	var err error
	m.Output.Body, err = json.Marshal(m.Output.BodyStruct)
	LogError(err)
	response, err = json.Marshal(m.Output)
	LogError(err)
	return
}

func initMessage() *Message {
	return &Message{
		Input:  InputMsg{},
		Output: OutputMsg{},
	}
}

func Dispatch(data string) *Message {
	msg := initMessage()
	msg.unpack(data)
	return msg
}
