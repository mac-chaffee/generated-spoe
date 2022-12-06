// This file is NOT generated.
// Consumers of this library should only use the types and methods in this file
// since the raw generated parser cannot preserve backwards compatibility
// as well as this hand-written file
package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

// Frame is a helper interface, mainly implemented by SPOPFrame
type Frame interface {
	NewFrameFromFile(string)
	NewFrameFromBytes([]byte)
}

// // FrameType tells you what kind of data to expect in this frame. Ack and Notify are the two important ones.
// type FrameType byte

// const (
// 	FrameTypeUnset             FrameType = 0
// 	FrameTypeHaproxyHello      FrameType = 1
// 	FrameTypeHaproxyDisconnect FrameType = 2
// 	FrameTypeNotify            FrameType = 3
// 	FrameTypeAgentHello        FrameType = 101
// 	FrameTypeAgentDisconnect   FrameType = 102
// 	FrameTypeAck               FrameType = 103
// )

// // ActionType is how you tell HAProxy whether to set or unset a set of vars
// type ActionType byte

// const (
// 	ActionTypeSetVar   byte = 1
// 	ActionTypeUnsetVar byte = 2
// )

// // VarScope is typically just set to VarScopeTransaction but otherwise unimportant
// type VarScope byte

// const (
// 	VarScopeProcess     byte = 1
// 	VarScopeSession     byte = 2
// 	VarScopeTransaction byte = 3
// 	VarScopeRequest     byte = 4
// 	VarScopeResponse    byte = 5
// )

// Action is how you tell HAProxy what to do inside ACK frames
type Action struct {
	ActionType Spop_ActionType
	Scope      Spop_VarScope
	Name       string
	Value      interface{}
}

// Message is how HAProxy sends info about a given request to you
type Message struct {
	Name string
	Args map[string]interface{}
}

// SPOPFrame represents the entire data portion of an SPOP TCP packet. Note: some fields are only set for specific message types.
type SPOPFrame struct {
	Frame
	FrameLen  uint32
	StreamId  int
	FrameId   int
	FrameType Spop_FrameType
	// Only applicable for FrameTypeNotify
	Messages []Message
	// Only applicable for FrameTypeAck
	Actions []Action
}

func NewFrameFromFile(filename string) (*SPOPFrame, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	stream := kaitai.NewStream(f)
	spop := NewSpop()
	if err = spop.Read(stream, nil, spop); err != nil {
		return nil, err
	}
	userFrame, err := fromRawFrame(spop)
	if err != nil {
		return nil, err
	}
	return userFrame, nil
}

func NewFrameFromBytes(b []byte) (*SPOPFrame, error) {
	buf := bytes.NewReader(b)
	stream := kaitai.NewStream(buf)
	spop := NewSpop()
	if err := spop.Read(stream, nil, spop); err != nil {
		return nil, err
	}
	userFrame, err := fromRawFrame(spop)
	if err != nil {
		return nil, err
	}
	return userFrame, nil
}

// fromRawFrame is the bulwark separating the messy generated parser from
// consumers of this library. SPOPFrame should never change, but the body of
// this function and the generated code can change as often as we need.
func fromRawFrame(spop *Spop) (*SPOPFrame, error) {
	// First set the easy-to-set fields
	frame := SPOPFrame{
		FrameLen:  spop.LenFrame,
		FrameType: spop.Frame.FrameType,
	}
	// Set metadata, which requires error handling
	streamId, err := spop.Frame.FrameMeta.StreamId.Value()
	if err != nil {
		return nil, err
	}
	frame.StreamId = streamId
	frameId, err := spop.Frame.FrameMeta.FrameId.Value()
	if err != nil {
		return nil, err
	}
	frame.FrameId = frameId
	// Determine if we need to set Args or Actions
	switch spop.Frame.FrameType {
	case Spop_FrameType__Ack:
		actions := spop.Frame.FramePayload.(*Spop_ListOfActions).Actions
		setActions(&frame, actions)
	case Spop_FrameType__Notify:
		messages := spop.Frame.FramePayload.(*Spop_ListOfMessages).Messages
		setMessages(&frame, messages)
	default:
		return nil, fmt.Errorf("Frame type %q is not implemented", spop.Frame.FrameType)
	}
	return &frame, nil
}

// Given a raw list of Spop_Action, converts them to the much-nicer Action type
func setActions(frame *SPOPFrame, actions []*Spop_Action) {
	frame.Actions = make([]Action, len(actions))
	for i, action := range actions {
		if action.ActionType == Spop_ActionType__SetVar {
			rawVar := action.ActionArgs.(*Spop_ActionSetVar)
			frame.Actions[i] = Action{
				ActionType: action.ActionType,
				Scope:      rawVar.VarScope,
				Name:       rawVar.VarName.StrData,
				Value:      rawVar.VarValue.TypeData,
			}
		} else if action.ActionType == Spop_ActionType__UnsetVar {
			rawVar := action.ActionArgs.(*Spop_ActionUnsetVar)
			frame.Actions[i] = Action{
				ActionType: action.ActionType,
				Scope:      rawVar.VarScope,
				Name:       rawVar.VarName.StrData,
				Value:      nil,
			}
		}
	}
}

// Given a raw list of Spop_Message, converts them to the much-nicer Message type
func setMessages(frame *SPOPFrame, messages []*Spop_Message) {
	frame.Messages = make([]Message, len(messages))
	for i, message := range messages {
		frame.Messages[i] = Message{
			Name: message.MessageName.StrData,
			Args: make(map[string]interface{}, message.NbArgs),
		}
		for _, arg := range message.Kvs {
			frame.Messages[i].Args[arg.KvName.StrData] = arg.KvValue
		}
	}
}
