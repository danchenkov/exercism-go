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
	c := make(chan rune, 256)
	d := make(chan int)
	for _, s := range ss {
		go func(st string) {
			d <- 1
			for _, r := range st {
				c <- r
			}
			d <- -1
		}(s)
	}

	go func() {
		var counter int
		for {
			counter += <-d
			if counter == 0 {
				close(c)
				return
			}
		}
	}()

	for symbol := range c {
		m[symbol]++
	}
	return m
}
