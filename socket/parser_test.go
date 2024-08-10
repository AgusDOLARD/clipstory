package socket

import "testing"

func TestParse(t *testing.T) {
	clips := []string{
		"socket://localhost:8080/clips/1",
		"socket://localhost:8080/clips/2",
		"socket://localhost:8080/clips/3",
	}

	b := MarshalPacket(clips)

	clipsGotten, err := UnmarshalPacket(b)
	if err != nil {
		t.Errorf("Error unmarshalling packet: %s", err)
	}

	if len(clips) != len(clipsGotten) {
		t.Errorf("Expected %d clips, got %d", len(clips), len(clipsGotten))
	}

	if clips[0] != clipsGotten[0] {
		t.Errorf("Expected %s, got %s", clips[0], clipsGotten[0])
	}

}
