package serializer_test

import (
	"fmt"
	"gRPC_study/pb"
	"gRPC_study/sample"
	"gRPC_study/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	fmt.Println("laptop1:", laptop1)
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	fmt.Println("laptop2:", laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

}

func TestName(t *testing.T) {
	fmt.Print(rand.Int())
}
