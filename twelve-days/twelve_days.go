package twelve

const testVersion = 1

const staPart = "On the "
const midPart = " day of Christmas my true love gave to me, "
const andPart = ", and "
const endPart = "."
const sepPart = ", "

var christmasDays = [...]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var christmasPresents = [...]string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var collectedPresents []string

func Song() (song string) {
	for day := 1; day <= len(christmasDays); day++ {
		song += Verse(day) + "\n"
	}
	return song
}

func Verse(day int) (verse string) {
	if day <= 1 {
		verse = christmasPresents[0]
	} else {
		verse = reverseJoin(christmasPresents[1:day]) + andPart + christmasPresents[0]
	}
	return staPart + christmasDays[day-1] + midPart + verse + endPart
}

func reverseJoin(input []string) (reverse string) {
	if len(input) == 0 {
		return
	}
	if len(input) == 1 {
		return input[0]
	}
	for i := len(input) - 1; i > 0; i-- {
		reverse += input[i] + sepPart
	}
	reverse += input[0]
	return reverse
}
