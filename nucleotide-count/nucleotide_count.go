package dna

import "fmt"

type Histogram map[byte]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for i := 0; i < len(d); i++ {
		switch d[i] {
		case 'A':
			h['A']++
		case 'C':
			h['C']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		default:
			return h, fmt.Errorf("Invalid nucleotide %s at position %d", string(d[i]), i)
		}
	}
	return h, nil
}
