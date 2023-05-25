package serializer_test

import (
	"gRPC_study/sample"
	"gRPC_study/serializer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"

	laptiop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptiop1, binaryFile)
	require.NoError(t, err)
}
