package packethandling

import (
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"

	"strings"
)

func Setup(ipAddr string, port int) (*net.UDPConn, error) {
	addr := net.UDPAddr{
		IP:   net.ParseIP(ipAddr),
		Port: port,
	}

	return net.ListenUDP("udp", &addr)
}

// Prints out a nice hex-view of the data
func FormatByteArray(data []byte, limit int) string {
	var sb strings.Builder
	const bytesPerLine = 16

	// Write header
	sb.WriteString("Offset    Hex                                        ASCII\n")
	sb.WriteString("--------  ----------------------------------------  ----------------\n")

	// Calculate how many bytes to actually process
	bytesToProcess := len(data)
	if limit > 0 && limit < bytesToProcess {
		bytesToProcess = limit
	}

	for i := 0; i < bytesToProcess; i += bytesPerLine {
		// Write offset
		sb.WriteString(fmt.Sprintf("%08X  ", i))

		// Write hex values
		for j := 0; j < bytesPerLine; j++ {
			if i+j < bytesToProcess {
				sb.WriteString(fmt.Sprintf("%02X ", data[i+j]))
			} else {
				sb.WriteString("   ")
			}
		}

		// Write ASCII representation
		sb.WriteString(" ")
		for j := 0; j < bytesPerLine && i+j < bytesToProcess; j++ {
			b := data[i+j]
			if b >= 32 && b <= 126 { // Printable ASCII range
				sb.WriteByte(b)
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	if bytesToProcess < len(data) {
		sb.WriteString(fmt.Sprintf("\n... %d more bytes not shown ...\n", len(data)-bytesToProcess))
	}

	return sb.String()
}

// Pretty prints the struct using spew
func FormatStruct(d ForzaHorizon5Packet) string {
	config := spew.ConfigState{
		Indent:                  "    ",
		MaxDepth:                0,
		DisableMethods:          true,
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}
	return config.Sdump(d)

}
