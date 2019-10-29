package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// TeamResult keeps current standings for a team
type TeamResult struct {
	matches int
	wins    int
	draws   int
	losses  int
}

func (t TeamResult) points() int {
	return t.wins*3 + t.draws
}

// Ranking is used for sorting
type Ranking struct {
	team   string
	result TeamResult
}

// Tally reads slice of strings with raw results; writes current tournament standing
func Tally(r io.Reader, w io.Writer) error {
	results := map[string]TeamResult{}
	var rankings []Ranking
	var a, b TeamResult
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
		if len(result[0]) == 0 {
			return fmt.Errorf("Invalid team name: %q", result[0])
		}
		if len(result[1]) == 0 {
			return fmt.Errorf("Invalid team name: %q", result[1])
		}

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

	for k, v := range results {
		rankings = append(rankings, Ranking{team: k, result: v})
	}
	sort.Slice(rankings, func(i, j int) bool {
		if rankings[i].result.points() == rankings[j].result.points() {
			return rankings[i].team < rankings[j].team
		}
		return rankings[i].result.points() > rankings[j].result.points()
	})

	fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")
	for _, r := range rankings {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n", r.team, r.result.matches, r.result.wins, r.result.draws, r.result.losses, r.result.points())
	}
	return nil
}
