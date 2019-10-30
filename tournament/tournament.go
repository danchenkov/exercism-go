package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// teamResult keeps current standings for a team
type teamResult struct {
	name    string
	matches int
	wins    int
	draws   int
	losses  int
}

func (t teamResult) points() int {
	return t.wins*3 + t.draws
}

// Tally reads slice of strings with raw results; writes current tournament standing
func Tally(r io.Reader, w io.Writer) error {
	results := map[string]teamResult{}
	var a, b teamResult
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Count(scanner.Text(), ";") != 2 {
			return fmt.Errorf("Invalid input: %q", scanner.Text())
		}

		result := strings.Split(scanner.Text(), ";")

		a = results[result[0]]
		b = results[result[1]]
		switch result[2] {
		case "win":
			a.matches++
			a.wins++
			b.matches++
			b.losses++
		case "loss":
			b.matches++
			b.wins++
			a.matches++
			a.losses++
		case "draw":
			a.matches++
			a.draws++
			b.matches++
			b.draws++
		default:
			return fmt.Errorf("Bad match result: %q", result[2])
		}
		results[result[0]] = a
		results[result[1]] = b
	}

	var teamResults []teamResult
	teamResults = make([]teamResult, len(results))
	var counter int
	for k, v := range results {
		v.name = k
		teamResults[counter] = v
		counter++
	}
	sort.Slice(teamResults, func(i, j int) bool {
		if teamResults[i].points() == teamResults[j].points() {
			return teamResults[i].name < teamResults[j].name
		}
		return teamResults[i].points() > teamResults[j].points()
	})

	fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")
	for _, r := range teamResults {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n", r.name, r.matches, r.wins, r.draws, r.losses, r.points())
	}
	return nil
}
