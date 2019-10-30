package main

import "fmt"

type Histogram map[byte]int

type DNA string

func (h Histogram) String() string {
	return fmt.Sprintf("A:%d C:%d G:%d T:%d", h['A'], h['C'], h['G'], h['T'])
}

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for i := 0; i < len(d); i++ {
		if _, ok := h[d[i]]; !ok {
			return h, fmt.Errorf("Invalid nucleotide %s at position %d (%s)", string(d[i]), i, d)
		}
		h[d[i]]++
	}
	return h, nil
}

func main() {
	var d DNA
	d = "AGACXXT"
	if h, e := d.Counts(); e != nil {
		fmt.Printf("Error: %s\n", e)
	} else {
		fmt.Println(h)
	}
}
