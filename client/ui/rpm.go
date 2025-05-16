package ui

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

const (
	meterWidth    = 50 // Total segments in the meter
	greenSegments = 30 // Number of green segments before color transition starts

)

func CreateRPMMeter() *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)
}

func UpdateRPMMeter(meter *tview.TextView, currentRPM, maxRPM float32) {
	if maxRPM == 0 {
		return
	}

	// Calculate how many segments should be filled
	percentage := currentRPM / maxRPM
	filledSegments := int(float32(meterWidth) * percentage)
	if filledSegments > meterWidth {
		filledSegments = meterWidth
	}

	var sb strings.Builder
	sb.WriteString("RPM: ")

	// Add the meter segments with color gradients
	for i := 0; i < meterWidth; i++ {
		if i < filledSegments {
			// Calculate color based on position
			if i < greenSegments {
				// Green segment
				sb.WriteString("[green]█")
			} else {
				// Transition from yellow to red
				position := float32(i-greenSegments) / float32(meterWidth-greenSegments)
				if position < 0.5 {
					sb.WriteString("[yellow]█")
				} else {
					sb.WriteString("[red]█")
				}
			}
		} else {
			// Empty segment
			sb.WriteString("[gray]░")
		}
	}

	// Add RPM value
	sb.WriteString(fmt.Sprintf("[-] %.0f/%.0f", currentRPM, maxRPM))

	meter.SetText(sb.String())
}
