package tournament

import (
	"bufio"
	"errors"
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

func (t *TeamResult) points() int {
	return t.wins*3 + t.draws
}

// Ranking is used for sorting
type Ranking struct {
	team   string
	result *TeamResult
}

func win(team1, team2 string, results map[string]*TeamResult) map[string]*TeamResult {
	if r, ok := results[team1]; ok {
		r.matches++
		r.wins++
		results[team1] = r
	} else {
		results[team1] = &TeamResult{matches: 1, wins: 1}
	}
	if r, ok := results[team2]; ok {
		r.matches++
		r.losses++
		results[team2] = r
	} else {
		results[team2] = &TeamResult{matches: 1, losses: 1}
	}
	return results
}

func draw(team1, team2 string, results map[string]*TeamResult) map[string]*TeamResult {
	if r, ok := results[team1]; ok {
		r.matches++
		r.draws++
		results[team1] = r
	} else {
		results[team1] = &TeamResult{matches: 1, draws: 1}
	}
	if r, ok := results[team2]; ok {
		r.matches++
		r.draws++
		results[team2] = r
	} else {
		results[team2] = &TeamResult{matches: 1, draws: 1}
	}
	return results
}

// Tally reads slice of strings with raw results; writes current tournament standing
func Tally(r io.Reader, w io.Writer) error {
	results := make(map[string]*TeamResult)
	var rankings []Ranking
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		if strings.Count(scanner.Text(), ";") != 2 {
			return errors.New("Invalid input: " + scanner.Text())
		}

		result := strings.Split(scanner.Text(), ";")
		if len(result[0]) == 0 || strings.HasPrefix(result[0], "#") || strings.HasPrefix(result[0], "//") {
			return errors.New("Invalid team name: " + result[0])
		}
		if len(result[1]) == 0 || strings.HasPrefix(result[1], "#") || strings.HasPrefix(result[1], "//") {
			return errors.New("Invalid team name: " + result[1])
		}

		switch result[2] {
		case "win":
			results = win(result[0], result[1], results)
		case "loss":
			results = win(result[1], result[0], results)
		case "draw":
			results = draw(result[0], result[1], results)
		default:
			return errors.New("Invalid result: " + result[2])
		}
	}

	for k, v := range results {
		r := new(Ranking)
		r.team = k
		r.result = v
		rankings = append(rankings, *r)
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
