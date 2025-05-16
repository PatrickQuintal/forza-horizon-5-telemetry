package main

import (
	"fmt"
	"forza-horizon-5-telemetry/client/ui"
	"forza-horizon-5-telemetry/shared/packettypes"
	"time"

	//"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"net"
)

const (
	meterWidth              = 50 // Total segments in the meter
	greenSegments           = 30 // Number of green segments before color transition starts
	defaultUpdatesPerSecond = 10 // Default update rate
)

func main() {
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

		addr := net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 9999,
		}

		conn, err := net.ListenUDP("udp", &addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		buf := make([]byte, 1024)
		var lastDash packettypes.ForzaHorizon5Packet

		for {

			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				app.QueueUpdateDraw(func() {
					textView.SetText(fmt.Sprintf("Read error: %v", err))
				})
				return
			}

			var dash packettypes.ForzaHorizon5Packet
			err = packettypes.ParsePacket(buf[:n], &dash)
			if err != nil {
				app.QueueUpdateDraw(func() {
					textView.SetText(fmt.Sprintf("Error parsing packet: %v", err))
				})
				return
			}

			// Store the latest data
			lastDash = dash

			// Wait for ticker before updating UI
			select {
			case <-ticker.C:

				app.QueueUpdateDraw(func() {
					if !isDebugView {
						ui.UpdateRPMMeter(rpmMeter, lastDash.GetCurrentEngineRpm(), lastDash.GetEngineMaxRpm())
						ui.UpdateSpeedometer(speedometer, lastDash.GetSpeedKMH())
						ui.UpdateLeftInfoPanel(leftInfoPanel, lastDash)
						ui.UpdateRightInfoPanel(rightInfoPanel, lastDash)
					} else {
						// Update debug view
						ui.UpdateDebugView(debugView, lastDash)
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
