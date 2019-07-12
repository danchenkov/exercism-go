package scale

import (
	"strings"
)

const testVersion = 1

const octave = 12
const (
	none = iota
	sharps
	flats
)

var scaleWithSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var scaleWithFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var intervals = map[string]int{"m": 1, "M": 2, "A": 3}

func pos(element string, array []string) int {
	for i, v := range array {
		if element == v {
			return i
		}
	}
	return -1
}

func Scale(tonic, interval string) []string {
	var scale int
	var currentIndex int
	var currentInterval string
	startingIndex := 0

	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = flats
		startingIndex = pos(strings.Title(tonic), scaleWithFlats)

	// case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#",
	default:
		scale = sharps
		startingIndex = pos(strings.Title(tonic), scaleWithSharps)
	}

	notes := []string{}
	currentIndex = startingIndex

	if len(interval) == 0 {
		interval = "mmmmmmmmmmmm"
	}

	for len(interval) > 0 {
		if scale == flats {
			notes = append(notes, scaleWithFlats[currentIndex])
		} else {
			notes = append(notes, scaleWithSharps[currentIndex])
		}

		currentInterval, interval = string(interval[0]), interval[1:]
		currentIndex = currentIndex + intervals[currentInterval]
		currentIndex = currentIndex % 12
	}

	if currentIndex != startingIndex {
		if scale == flats {
			notes = append(notes, scaleWithFlats[currentIndex])
		} else {
			notes = append(notes, scaleWithSharps[currentIndex])
		}
	}

	return notes
}
