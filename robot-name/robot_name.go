package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const testVersion = 1

var max = 26 * 26 * 1000
var names = map[string]bool{"": true}

// Robot properties
type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name returns a random robot name
func (robot *Robot) Name() (string, error) {
	var err error

	if robot.name != "" {
		return robot.name, err
	}

	if len(names) > max {
		return "", errors.New("Namespace exhausted")
	}

	for names[robot.name] {
		robot.name = fmt.Sprintf("%c%c%00d", rand.Intn(26)+'A', rand.Intn(26)+'A', rand.Intn(1000))
	}
	names[robot.name] = true

	return robot.name, nil
}

// Reset clears out namespace for robots
func (robot *Robot) Reset() {
	robot.name = ""
}
