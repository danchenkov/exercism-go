package robotname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const testVersion = 1

var names map[string]struct{}

type Robot struct {
	name string
}

func init() {
	names = make(map[string]struct{})
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var b strings.Builder
	var name string
	for ok := true; ok; _, ok = names[name] {
		for i := 0; i < 2; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		name = b.String() + fmt.Sprintf("%03d", rand.Intn(1000))
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
