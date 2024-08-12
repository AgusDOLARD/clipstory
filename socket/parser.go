package socket

import (
	"encoding/binary"
	"fmt"
	"slices"
	"strings"
)

const (
	VERSION       = 2
	HEADER_LENGTH = 3 // 1 byte for version + 2 bytes for length
)

// MarshalPacket takes a slice of clipboard entries and returns a byte slice
// representing the packet.
func MarshalPacket(clips []string) []byte {
	squashed := strings.Join(clips, "\r\n")
	length := uint16(len(squashed))
	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, length)

	// Create a byte slice with the appropriate capacity
	b := make([]byte, 0, HEADER_LENGTH+len(squashed))
	b = append(b, VERSION)
	b = append(b, lengthBytes...)
	b = append(b, []byte(squashed)...)
	return b
}

// UnmarshalPacket takes a byte slice and returns the corresponding slice of clipboard entries.
func UnmarshalPacket(b []byte) ([]string, error) {
	if len(b) < HEADER_LENGTH {
		return nil, fmt.Errorf("packet too short: minimum length is %d, got %d", HEADER_LENGTH, len(b))
	}

	version := b[0]
	if version != VERSION {
		return nil, fmt.Errorf("invalid version: expected %d, got %d", VERSION, version)
	}

	length := binary.BigEndian.Uint16(b[1:3]) // Read the 2-byte length
	if len(b) < HEADER_LENGTH+int(length) {
		return nil, fmt.Errorf("packet too short for the specified length: expected %d, got %d", HEADER_LENGTH+length, len(b))
	}

	data := b[HEADER_LENGTH : HEADER_LENGTH+int(length)]
	clips := strings.Split(string(data), "\r\n")
	slices.Reverse(clips)
	return clips, nil
}
