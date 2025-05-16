package debugstreamreader

import (
	"fmt"
	"os"
)

type DebugStreamReader struct {
	data       []byte
	position   int
	packetSize int
}

func NewDebugStreamReader(filepath string, packetSize int) (*DebugStreamReader, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Validate file size is multiple of packet size
	if len(data)%packetSize != 0 {
		return nil, fmt.Errorf("file size %d is not a multiple of packet size %d", len(data), packetSize)
	}

	return &DebugStreamReader{
		data:       data,
		position:   0,
		packetSize: packetSize,
	}, nil
}

func (r *DebugStreamReader) ReadNext() ([]byte, error) {
	// Check if we've reached the end of the data
	if r.position >= len(r.data) {
		r.position = 0 // Loop back to start
	}

	// Validate packet boundary
	if (len(r.data) - r.position) < r.packetSize {
		return nil, fmt.Errorf("incomplete packet at position %d", r.position)
	}

	// Extract exactly one packet
	packet := make([]byte, r.packetSize)
	copy(packet, r.data[r.position:r.position+r.packetSize])
	r.position += r.packetSize

	return packet, nil
}
