package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can not marshal proto message to binary: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 8644)
	if err != nil {
		return fmt.Errorf("can not write data to file: %w", err)
	}

	return nil
}
