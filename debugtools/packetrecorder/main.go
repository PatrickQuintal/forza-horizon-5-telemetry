package main

import (
	"forza-horizon-5-telemetry/shared/packethandling"
	"log"
	"os"
)

func main() {
	const (
		packetsPerSecond = 60
		secondsToRecord  = 10
		packetSize       = 324
		totalPackets     = packetsPerSecond * secondsToRecord
	)

	// Setup the UDP connection
	conn, err := packethandling.Setup("127.0.0.1", 9999)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Pre-allocate with exact size needed, but length 0
	filebuf := make([]byte, 0, totalPackets*packetSize)
	buf := make([]byte, packetSize) // Only need packetSize bytes

	for i := 0; i < totalPackets; i++ {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n != packetSize {
			log.Printf("Warning: Packet %d: Expected %d bytes, got %d bytes\n", i, packetSize, n)
			continue
		}

		var fh5Packet packethandling.ForzaHorizon5Packet
		err = packethandling.ParsePacket(buf[:n], &fh5Packet)
		if err != nil {
			log.Printf("Warning: Packet %d: Parse error: %v\n", i, err)
			continue
		}

		filebuf = append(filebuf, buf[:n]...)

		if i%packetsPerSecond == 0 {
			log.Printf("Recording: %d/%d seconds\n", i/packetsPerSecond, secondsToRecord)
		}
	}

	expectedSize := totalPackets * packetSize
	if len(filebuf) != expectedSize {
		log.Printf("Warning: Expected file size %d, got %d\n", expectedSize, len(filebuf))
	}

	err = os.WriteFile("./debugstream", filebuf, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File written successfully: %d packets (%d bytes)\n", len(filebuf)/packetSize, len(filebuf))
}
