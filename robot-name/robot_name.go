package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const testVersion = 1

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var names map[string]struct{}

type Robot struct {
	name string
}

func init() {
	names = make(map[string]struct{})
}

func randStringBytesMaskImprSrc(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func generateName() string {
	var name string
	for ok := true; ok; _, ok = names[name] {
		name = randStringBytesMaskImprSrc(2) + fmt.Sprintf("%03d", rand.Intn(1000))
	}
	names[name] = struct{}{}
	return name
}

func (robot *Robot) Name() (string, error) {
	if len(robot.name) == 0 {
		robot.name = generateName()
	}
	return robot.name, nil
}

func (robot *Robot) Reset() (string, error) {
	robot.name = generateName()
	return robot.name, nil
}
