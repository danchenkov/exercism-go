package house

import "fmt"

const testVersion = 1
const first = "This is "
const that = "that"

var subjects = []string{
	"the horse and the hound and the horn",
	"the farmer sowing his corn",
	"the rooster that crowed in the morn",
	"the priest all shaven and shorn",
	"the man all tattered and torn",
	"the maiden all forlorn",
	"the cow with the crumpled horn",
	"the dog",
	"the cat",
	"the rat",
	"the malt",
	"the house that Jack built.",
}
var actions = []string{
	"belonged to",
	"kept",
	"woke",
	"married",
	"kissed",
	"milked",
	"tossed",
	"worried",
	"killed",
	"ate",
	"lay in",
	"",
}
var length = len(subjects)

func Song() (song string) {
	for i := 1; i < length; i++ {
		song += Verse(i) + "\n\n"
	}
	song += Verse(length)
	return song
}

func Verse(n int) (verse string) {
	verse = first
	for i := length - n; i < length; i++ {
		verse += fmt.Sprintf("%s", subjects[i])
		if i < length-1 {
			verse += fmt.Sprintf("\n%s %s ", that, actions[i])
		}
	}
	return verse
}
