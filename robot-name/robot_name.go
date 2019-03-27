package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const testVersion = 1

var max = 676000
var names map[string]bool

// Robot properties
type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
	names = make(map[string]bool)
}

// Name returns a random robot name
func (robot *Robot) Name() (string, error) {
	var err error
	if len(robot.name) == 0 {
		if len(names) >= max {
			return "", errors.New("Namespace exhausted")
		}

		robot.name = fmt.Sprintf("%c%c%d", rand.Intn(26)+'A', rand.Intn(26)+'A', rand.Intn(1000))
		for names[robot.name] {
			robot.name = fmt.Sprintf("%c%c%d", rand.Intn(26)+'A', rand.Intn(26)+'A', rand.Intn(1000))
		}
		names[robot.name] = true

		return robot.name, nil
	}
	return robot.name, err
}

// Reset clears out namespace for robots
func (robot *Robot) Reset() (string, error) {
	robot.name = ""
	return robot.name, nil
}
