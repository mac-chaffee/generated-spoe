package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

// parse is a helper function to parse a test binary file
func parse(filename string) *Spop {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stream := kaitai.NewStream(f)
	spop := NewSpop()
	if err = spop.Read(stream, nil, spop); err != nil {
		panic(err)
	}
	return spop
}

// Reference resources/spop-example.txt to see the frame decoded manually
func TestAckFrame(t *testing.T) {
	spop := parse("../../resources/ack_frame.bin")
	// Must use EqualValues since LenFrame is actually uint32=0x14
	assert.EqualValues(t, 20, spop.LenFrame)
	assert.Equal(t, Spop_FrameType__Ack, spop.Frame.FrameType)

	meta := spop.Frame.FrameMeta
	assert.Equal(t, true, meta.MetaFlags.FinFlag)
	assert.Equal(t, false, meta.MetaFlags.AbortFlag)
	stream_id, err := meta.StreamId.Value()
	assert.NoError(t, err)
	assert.Equal(t, 2724629, stream_id)
	frame_id, err := meta.FrameId.Value()
	assert.NoError(t, err)
	assert.Equal(t, 1, frame_id)

	actions := spop.Frame.FramePayload.(*Spop_ListOfActions).Actions
	assert.Equal(t, 1, len(actions))
	assert.Equal(t, Spop_ActionType__SetVar, actions[0].ActionType)
	assert.EqualValues(t, 3, actions[0].NbArgs)

	args := actions[0].ActionArgs.(*Spop_ActionSetVar)
	assert.Equal(t, Spop_VarScope__Transaction, args.VarScope)
	assert.Equal(t, "fail", args.VarName.StrData)

	var_value := args.VarValue.TypeData.(*Varint)
	val, err := var_value.Value()
	assert.NoError(t, err)
	assert.Equal(t, 0, val)
}

func TestNotifyFrame(t *testing.T) {
	spop := parse("../../resources/notify_frame.bin")
	// Must use EqualValues since LenFrame is actually uint32
	assert.EqualValues(t, 883, spop.LenFrame)
	assert.Equal(t, Spop_FrameType__Notify, spop.Frame.FrameType)

	meta := spop.Frame.FrameMeta
	assert.Equal(t, true, meta.MetaFlags.FinFlag)
	assert.Equal(t, false, meta.MetaFlags.AbortFlag)
	stream_id, err := meta.StreamId.Value()
	assert.NoError(t, err)
	assert.EqualValues(t, 2724629, stream_id)
	frame_id, err := meta.FrameId.Value()
	assert.NoError(t, err)
	assert.Equal(t, 1, frame_id)

	messages := spop.Frame.FramePayload.(*Spop_ListOfMessages).Messages
	assert.Equal(t, 1, len(messages))
	assert.EqualValues(t, 12, messages[0].NbArgs)

	args := messages[0].Kvs
	assert.Equal(t, 12, len(args))
	assert.Equal(t, "app", args[0].KvName.StrData)
	assert.Equal(t, "containers-dev.apps.renci.org", args[0].KvValue.TypeData.(*Spop_SpopString).StrData)
	assert.Equal(t, "id", args[1].KvName.StrData)
	assert.Equal(t, "705e62b0-765a-4494-bb71-29c8f35ac387", args[1].KvValue.TypeData.(*Spop_SpopString).StrData)
	// ...
	assert.Equal(t, "headers", args[10].KvName.StrData)
	header_len, err := args[10].KvValue.TypeData.(*Spop_SpopString).StrLen.Value()
	assert.NoError(t, err)
	assert.EqualValues(t, 665, header_len)
}
