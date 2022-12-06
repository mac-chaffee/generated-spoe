package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper(t *testing.T) {
	filename := "../../resources/ack_frame.bin"
	frame, err := NewFrameFromFile(filename)
	assert.NoError(t, err)
	rawFrame := parse(filename)

	assert.EqualValues(t, rawFrame.LenFrame, frame.FrameLen)
	assert.Equal(t, rawFrame.Frame.FrameType, frame.FrameType)
	streamId, err := rawFrame.Frame.FrameMeta.StreamId.Value()
	assert.NoError(t, err)
	assert.Equal(t, streamId, frame.StreamId)
	frameId, err := rawFrame.Frame.FrameMeta.FrameId.Value()
	assert.NoError(t, err)
	assert.Equal(t, frameId, frame.FrameId)
	action := rawFrame.Frame.FramePayload.(*Spop_ListOfActions).Actions[0].ActionArgs.(*Spop_ActionSetVar)
	assert.Equal(t, action.VarName.StrData, frame.Actions[0].Name)
	rawVal, err := action.VarValue.TypeData.(*Varint).Value()
	assert.NoError(t, err)
	val, err := frame.Actions[0].Value.(*Varint).Value()
	assert.NoError(t, err)
	assert.Equal(t, rawVal, val)
}
