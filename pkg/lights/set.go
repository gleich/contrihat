package lights

import (
	"time"

	"github.com/Matt-Gleich/contrihat/pkg/api"
	"github.com/Matt-Gleich/contrihat/pkg/config"
	"github.com/Matt-Gleich/logoru"
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
	"gopkg.in/go-playground/colors.v1"
)

// Set the lights on the sense hat
func Set(contributions api.Query, firstRun bool, configuration config.Outline) {
	// Setting color levels
	colorMap := map[string]string{}
	colors := contributions.Viewer.ContributionsCollection.ContributionCalendar.Colors
	if len(configuration.Levels) == 4 {
		for i, level := range configuration.Levels {
			colorMap[colors[i]] = level
		}
	} else {
		for _, color := range colors {
			if contributions.Viewer.ContributionsCollection.ContributionCalendar.IsHalloween {
				switch color {
				case "var(--color-calendar-halloween-graph-day-L1-bg)":
					colorMap[color] = "#FFEE4A"
				case "var(--color-calendar-halloween-graph-day-L2-bg)":
					colorMap[color] = "#FFC501"
				case "var(--color-calendar-halloween-graph-day-L3-bg)":
					colorMap[color] = "#FE9600"
				case "var(--color-calendar-halloween-graph-day-L4-bg)":
					colorMap[color] = "#03001C"
				default:
					colorMap[color] = "#000000"
				}
			} else {
				switch color {
				case "var(--color-calendar-graph-day-L1-bg)":
					colorMap[color] = "#9BE9A8"
				case "var(--color-calendar-graph-day-L2-bg)":
					colorMap[color] = "#40C463"
				case "var(--color-calendar-graph-day-L3-bg)":
					colorMap[color] = "#31A14E"
				case "var(--color-calendar-graph-day-L4-bg)":
					colorMap[color] = "#216E38"
				}
			}
		}
	}

	fb := screen.NewFrameBuffer()
	days := mergeDays(contributions)
	var (
		x int
		y int
	)
	for _, day := range days {
		if x == 8 {
			y++
			x = 0
		}
		if y == 8 {
			break
		}
		fb.SetPixel(x, y, convert(colorMap[day]))
		if firstRun {
			time.Sleep(50 * time.Millisecond)
		}
		err := screen.Draw(fb)
		if err != nil {
			logoru.Error("Failed to draw screen;", err)
		}
		x++
	}
	logoru.Success("Updated lights!")
}

// Merging all weeks
func mergeDays(contributions api.Query) (days []string) {
	for _, week := range contributions.Viewer.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			days = append([]string{day.Color}, days...)
		}
	}
	return days
}

// Convert the hex code to bobblehat color
func convert(rawHex string) color.Color {
	hex, err := colors.ParseHEX(rawHex)
	if err != nil {
		logoru.Warning("Failed to parse hex code;", hex)
		return color.New(0, 0, 0)
	}
	rgb := hex.ToRGB()
	return color.New(rgb.R, rgb.G, rgb.B)
}
