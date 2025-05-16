package ui

import (
	"fmt"
	"forza-horizon-5-telemetry/shared/packettypes"
	"strings"

	"github.com/rivo/tview"
)

func CreateInfoPanel() *tview.TextView {
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)
	tv.SetBorder(true)
	return tv
}

func UpdateLeftInfoPanel(panel *tview.TextView, dash packettypes.ForzaHorizon5Packet) {
	var sb strings.Builder

	sb.WriteString("[yellow]Car Information[white]\n")
	sb.WriteString("---------------\n")
	sb.WriteString(fmt.Sprintf("Speed: %.0f km/h\n", dash.GetSpeedKMH()))
	sb.WriteString(fmt.Sprintf("Power: %.0f hp\n", dash.GetPower()))
	sb.WriteString(fmt.Sprintf("Torque: %.0f Nm\n", dash.GetTorque()))
	sb.WriteString(fmt.Sprintf("Car Class: %d\n", dash.GetCarClass()))
	sb.WriteString(fmt.Sprintf("PI: %d\n", dash.GetCarPerformanceIndex()))

	panel.SetText(sb.String())
}

func UpdateRightInfoPanel(panel *tview.TextView, dash packettypes.ForzaHorizon5Packet) {
	var sb strings.Builder

	sb.WriteString("[yellow]Race Information[white]\n")
	sb.WriteString("----------------\n")
	sb.WriteString(fmt.Sprintf("Position: %d\n", dash.GetRacePosition()))
	sb.WriteString(fmt.Sprintf("Lap: %d\n", dash.GetLapNumber()))
	sb.WriteString(fmt.Sprintf("Best Lap: %.2f\n", dash.BestLap))
	sb.WriteString(fmt.Sprintf("Last Lap: %.2f\n", dash.LastLap))
	sb.WriteString(fmt.Sprintf("Current Lap: %.2f\n", dash.CurrentLap))

	panel.SetText(sb.String())
}
