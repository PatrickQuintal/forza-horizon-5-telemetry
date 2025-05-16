package main

import (
	"flag"
	"fmt"
	"forza-horizon-5-telemetry/client/ui"
	"forza-horizon-5-telemetry/debugtools/debugstreamreader"
	"forza-horizon-5-telemetry/shared/packethandling"
	"time"

	//"github.com/gdamore/tcell/v2"
	"log"

	"github.com/rivo/tview"
)

const (
	defaultUpdatesPerSecond = 10 // Default update rate
)

func main() {
	debugMode := flag.Bool("debug", false, "Use debug stream instead of UDP")
	debugFile := flag.String("debugfile", "debugstream", "Path to debug stream file")
	flag.Parse()

	app := tview.NewApplication()

	// Create main flex container (vertical)
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Create normal view
	normalView := tview.NewFlex().SetDirection(tview.FlexRow)

	// Create horizontal flex for top panels
	topFlex := tview.NewFlex().SetDirection(tview.FlexColumn)

	// Create info panels
	leftInfoPanel := ui.CreateInfoPanel()
	rightInfoPanel := ui.CreateInfoPanel()

	// Add panels to top flex with equal weight
	topFlex.AddItem(leftInfoPanel, 0, 1, false)
	topFlex.AddItem(rightInfoPanel, 0, 1, false)

	// Create bottom panel for meters
	bottomFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	rpmMeter := ui.CreateRPMMeter()
	speedometer := ui.CreateSpeedometer()

	// Add bottom flex with fixed heights
	bottomFlex.AddItem(rpmMeter, 3, 0, false)    // Fixed 3 lines for RPM meter
	bottomFlex.AddItem(speedometer, 7, 0, false) // Fixed 7 lines for speedometer

	// Add both flexboxes to main container with fixed heights
	normalView.AddItem(topFlex, 8, 0, false)     // Fixed 8 lines for info panels
	normalView.AddItem(bottomFlex, 10, 0, false) // Fixed 10 lines for meters

	// Create debug view (modify this part)
	debugView := ui.CreateDebugView()

	// Create toggle button and its function
	toggleButton := tview.NewButton("Toggle Debug View")

	isDebugView := false

	// Modify toggle button function
	toggleButton.SetSelectedFunc(func() {
		isDebugView = !isDebugView
		toggleButton.SetLabel("Toggle " + map[bool]string{true: "Normal View", false: "Debug View"}[isDebugView])
		if !isDebugView {
			mainFlex.RemoveItem(debugView)
			mainFlex.AddItem(normalView, 0, 1, true)
		} else {
			mainFlex.RemoveItem(normalView)
			mainFlex.AddItem(debugView, 0, 1, true)

		}
	})

	// Add button to main flex at top
	mainFlex.AddItem(toggleButton, 1, 0, false)
	// Add normal view as default
	mainFlex.AddItem(normalView, 0, 1, true)

	textView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetWrap(true).
		SetScrollable(true).
		SetText("Waiting for data...").
		SetChangedFunc(func() {
			app.Draw()
		})

	// Create ticker for rate limiting
	updatesPerSecond := defaultUpdatesPerSecond
	ticker := time.NewTicker(time.Second / time.Duration(updatesPerSecond))
	defer ticker.Stop()

	go func() {
		var readPacket func() ([]byte, error)

		if *debugMode {
			reader, err := debugstreamreader.NewDebugStreamReader(*debugFile, 324)
			if err != nil {
				log.Fatal(err)
			}
			readPacket = reader.ReadNext
		} else {

			conn, err := packethandling.Setup("127.0.0.1", 9999)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			buf := make([]byte, 1024)
			readPacket = func() ([]byte, error) {
				n, _, err := conn.ReadFromUDP(buf)
				if err != nil {
					return nil, err
				}
				return buf[:n], nil
			}

		}
		var fh5Packet packethandling.ForzaHorizon5Packet

		for {

			data, err := readPacket()
			if err != nil {
				app.QueueUpdateDraw(func() {
					textView.SetText(fmt.Sprintf("Read error: %v", err))
				})
				return
			}

			err = packethandling.ParsePacket(data, &fh5Packet)
			if err != nil {
				app.QueueUpdateDraw(func() {
					textView.SetText(fmt.Sprintf("Error parsing packet: %v", err))
				})
				return
			}

			// Wait for ticker before updating UI
			select {
			case <-ticker.C:

				app.QueueUpdateDraw(func() {
					if !isDebugView {
						ui.UpdateRPMMeter(rpmMeter, fh5Packet.GetCurrentEngineRpm(), fh5Packet.GetEngineMaxRpm())
						ui.UpdateSpeedometer(speedometer, fh5Packet.GetSpeedKMH())
						ui.UpdateLeftInfoPanel(leftInfoPanel, fh5Packet)
						ui.UpdateRightInfoPanel(rightInfoPanel, fh5Packet)
					} else {
						// Update debug view
						ui.UpdateDebugView(debugView, fh5Packet)
					}
				})

			default:
				// No tick yet
				continue
			}

		}
	}()

	if err := app.SetRoot(mainFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
