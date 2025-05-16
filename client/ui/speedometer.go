package ui

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

const (
	speedometerWidth = 50  // Increased width for better detail
	maxSpeedDisplay  = 400 // Maximum speed in km/h
	speedSegments    = 8   // Number of major segments
)

func CreateSpeedometer() *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
}

func UpdateSpeedometer(meter *tview.TextView, speedKMH float32) {
	var sb strings.Builder

	// Create more detailed arc shape
	arc := []string{
		"      [white]╭" + strings.Repeat("─", speedometerWidth-2) + "╮",
		"   [white]╭╯" + strings.Repeat(" ", speedometerWidth-2) + "╰╮",
		"[white]╭╯  " + strings.Repeat(" ", speedometerWidth-6) + "  ╰╮",
	}

	// Add the arc
	for _, line := range arc {
		sb.WriteString(line + "\n")
	}

	// Add digital speed with larger format
	sb.WriteString(fmt.Sprintf("[yellow]%3.0f[white]km/h\n", speedKMH))

	// Create speed markers with labels
	markerLine := createSpeedMarkers(speedKMH)
	sb.WriteString(markerLine + "\n")

	// Add speed labels
	speedLabels := createSpeedLabels()
	sb.WriteString(speedLabels)

	meter.SetText(sb.String())
}

func createSpeedMarkers(speed float32) string {
	var sb strings.Builder
	percentage := speed / maxSpeedDisplay
	needlePos := int(percentage * float32(speedometerWidth))

	sb.WriteString("[white]│")

	for i := 0; i < speedometerWidth; i++ {
		isMajorTick := i%(speedometerWidth/speedSegments) == 0
		isMinorTick := i%(speedometerWidth/speedSegments/2) == 0

		switch {
		case i == needlePos:
			// Make needle more visible
			if speed < maxSpeedDisplay*0.3 {
				sb.WriteString("[green]█")
			} else if speed < maxSpeedDisplay*0.7 {
				sb.WriteString("[yellow]█")
			} else {
				sb.WriteString("[red]█")
			}
		case i < needlePos:
			position := float32(i) / float32(speedometerWidth)
			if isMajorTick {
				if position < 0.3 {
					sb.WriteString("[green]┃")
				} else if position < 0.7 {
					sb.WriteString("[yellow]┃")
				} else {
					sb.WriteString("[red]┃")
				}
			} else if isMinorTick {
				if position < 0.3 {
					sb.WriteString("[green]┊")
				} else if position < 0.7 {
					sb.WriteString("[yellow]┊")
				} else {
					sb.WriteString("[red]┊")
				}
			} else {
				if position < 0.3 {
					sb.WriteString("[green]━")
				} else if position < 0.7 {
					sb.WriteString("[yellow]━")
				} else {
					sb.WriteString("[red]━")
				}
			}
		default:
			if isMajorTick {
				sb.WriteString("[gray]┃")
			} else if isMinorTick {
				sb.WriteString("[gray]┊")
			} else {
				sb.WriteString("[gray]━")
			}
		}
	}
	sb.WriteString("[white]│")
	return sb.String()
}

func createSpeedLabels() string {
	var sb strings.Builder
	sb.WriteString(strings.Repeat(" ", 3))

	for i := 0; i <= speedSegments; i++ {
		speed := (i * maxSpeedDisplay) / speedSegments
		if i == 0 {
			sb.WriteString(fmt.Sprintf("[gray]%-6d", speed))
		} else if i == speedSegments {
			sb.WriteString(fmt.Sprintf("[gray]%d", speed))
		} else {
			sb.WriteString(fmt.Sprintf("[gray]%-5d", speed))
		}
	}
	return sb.String()
}
