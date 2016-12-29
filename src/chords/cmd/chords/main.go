package main

import (
	"fmt"
	"os"
)

type chord struct {
	Name  string
	Notes []string
}

var intervals = map[[10]int]string{
	[10]int{0, 4, 3}:    "Major",
	[10]int{0, 3, 4}:    "Minor",
	[10]int{0, 4, 3, 3}: "Domininant seventh",
	[10]int{0, 4, 3, 4}: "Major seventh",
}

var notes = map[string]int{
	"C":  0,
	"C#": 1,
	"D":  2,
	"D#": 3,
	"E":  4,
	"F":  5,
	"F#": 6,
	"G":  7,
	"G#": 8,
	"A":  9,
	"A#": 10,
	"B":  11,
}

func chorder(n []string) (ch chord, err error) {
	var intervalNotes [10]int

	for i, v := range n {
		if i == 0 {
			continue
		}
		if notes[v] > notes[n[i-1]] {
			intervalNotes[i] = notes[v] - notes[n[i-1]]
		} else {
			intervalNotes[i] = notes[v] + 12 - notes[n[i-1]]
		}
	}

	chordInterval, found := intervals[intervalNotes]
	if !found {
		return chord{}, fmt.Errorf("chords: no chord found for interval %v", intervalNotes)
	}

	ch = chord{chordInterval, n}

	return ch, err
}

func main() {
	chord, err := chorder([]string{"C", "E", "G"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(chord)
	}
}
