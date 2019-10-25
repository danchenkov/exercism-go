package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamResult struct {
	matches int
	wins    int
	draws   int
	losses  int
}

func (t *TeamResult) points() int {
	return t.wins*3 + t.draws
}

type Ranking struct {
	team   string
	result *TeamResult
}

type byRanking []Ranking

func win(team1, team2 string, results map[string]*TeamResult) map[string]*TeamResult {
	// log.Print("win: " + team1)
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
	// log.Print("draw: " + team1 + "/" + team2)
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

func Tally(r io.Reader, w io.Writer) error {
	results := make(map[string]*TeamResult)
	var rankings []Ranking
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			// log.Print("Empty string in the input; ignoring")
			continue
		}

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			// log.Print("Comment in the input; ignoring")
			continue
		}

		if strings.Count(scanner.Text(), ";") != 2 {
			return errors.New("Invalid input: " + scanner.Text())
		}
		// log.Print(line)

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
	sort.Sort(byRanking(rankings))

	fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")
	for _, r := range rankings {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n", r.team, r.result.matches, r.result.wins, r.result.draws, r.result.losses, r.result.points())
	}
	return nil
}

func (r byRanking) Len() int {
	return len(r)
}

func (r byRanking) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRanking) Less(i, j int) bool {
	if r[i].result.points() == r[j].result.points() {
		if r[i].result.wins == r[j].result.wins {
			if r[i].result.matches == r[j].result.matches {
				if r[i].result.losses == r[j].result.losses {
					return r[i].team < r[j].team
				}
				return r[i].result.losses < r[j].result.losses
			}
			return r[i].result.matches > r[j].result.matches
		}
		return r[i].result.wins > r[j].result.wins
	}
	return r[i].result.points() > r[j].result.points()
}

// func main() {
// 	blob := "\n\nAllegoric Alaskians;Blithering Badgers;win\n##Allegoric Alaskians;Blithering Badgers;win\nDevastating Donkeys;Courageous Californians;draw\nDevastating Donkeys;Allegoric Alaskians;win\nCourageous Californians;Blithering Badgers;loss\nBlithering Badgers;Devastating Donkeys;loss\nAllegoric Alaskians;Courageous Californians;win"
// 	buf := bytes.NewBufferString(blob)
// 	if err := Tally(buf, os.Stdout); err != nil {
// 		fmt.Printf("Error: %s", err)
// 	}
// }
