package socket

import (
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	VERSION       = 1
	HEADER_LENGTH = 3 // 1 byte for version + 2 bytes for length
)

func MarshalPacket(clips []string) []byte {
	squashed := strings.Join(clips, "\r\n")
	length := uint16(len(squashed))
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, length)

	b := make([]byte, 0, HEADER_LENGTH+len(squashed))
	b = append(b, VERSION)
	b = append(b, lengthBytes...)
	b = append(b, []byte(squashed)...)
	return b
}

func UnmarshalPacket(b []byte) ([]string, error) {
	if len(b) < HEADER_LENGTH {
		return nil, fmt.Errorf("packet too short")
	}

	version := b[0]
	if version != VERSION {
		return nil, fmt.Errorf("invalid version: %d", version)
	}

	length := binary.BigEndian.Uint16(b[1:3]) // Read the 2-byte length
	if len(b) < HEADER_LENGTH+int(length) {
		return nil, fmt.Errorf("packet too short for the specified length")
	}

	data := b[HEADER_LENGTH : HEADER_LENGTH+int(length)]
	clips := strings.Split(string(data), "\r\n")
	return clips, nil
}
