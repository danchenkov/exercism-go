package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap)
	for _, s := range ss {
		go func(st string) {
			c <- Frequency(st)
		}(s)
	}
	for range ss {
		t := <-c
		for i, v := range t {
			m[i] += v
		}
	}

	return m
}
