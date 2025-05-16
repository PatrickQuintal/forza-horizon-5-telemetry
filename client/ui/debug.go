package ui

import (
	"forza-horizon-5-telemetry/shared/packettypes"
	"github.com/rivo/tview"
)

func CreateDebugView() *tview.TextView {
	debugView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetScrollable(true)

	debugView.SetBorder(true).
		SetTitle("Debug View")

	return debugView
}

func UpdateDebugView(debugView *tview.TextView, data packettypes.ForzaHorizon5Packet) {
	debugView.SetText(packettypes.FormatStruct(data))
}
