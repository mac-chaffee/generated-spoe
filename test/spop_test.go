package test

import (
	"os"
	"testing"

	"github.com/mac-chaffee/generated-spoe/parser"
	"github.com/stretchr/testify/assert"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

// parse is a helper function to parse a test binary file
func parse(filename string) *parser.Spop {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stream := kaitai.NewStream(f)
	spop := parser.NewSpop()
	if err = spop.Read(stream, nil, spop); err != nil {
		panic(err)
	}
	return spop
}

// Reference resources/spop-example.txt to see the frame decoded manually
func TestAckFrame(t *testing.T) {
	spop := parse("../resources/ack_frame.bin")
	// Must use EqualValues since FrameLen is actually uint32=0x14
	assert.EqualValues(t, 20, spop.FrameLen)
	assert.Equal(t, parser.Spop_FrameType__Ack, spop.FrameType)

	meta := spop.FrameMeta
	assert.Equal(t, true, meta.MetaFlags.FinFlag)
	assert.Equal(t, false, meta.MetaFlags.AbortFlag)
	// TODO: Uncomment after we fix varint_test
	stream_id, err := meta.StreamId.Value()
	assert.NoError(t, err)
	assert.Equal(t, 2724629, stream_id)
	frame_id, err := meta.FrameId.Value()
	assert.NoError(t, err)
	assert.Equal(t, 1, frame_id)

	payload := spop.FramePayload.(*parser.Spop_ListOfActions)
	assert.Equal(t, parser.Spop_ActionType__SetVar, payload.ActionType)
	assert.EqualValues(t, 3, payload.NbArgs)

	args := payload.ActionArgs.(*parser.Spop_ActionSetVar)
	assert.Equal(t, parser.Spop_VarScope__Transaction, args.VarScope)
	assert.Equal(t, "fail", args.VarName.StrData)

	var_value := args.VarValue.TypeData.(*parser.Varint)
	val, err := var_value.Value()
	assert.NoError(t, err)
	assert.Equal(t, 0, val)
}
