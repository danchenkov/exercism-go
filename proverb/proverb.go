package proverb

func Proverb(rhyme []string) []string {
	message := []string{}

	for i := 0; i < len(rhyme); i++ {
		if i == len(rhyme)-1 {
			message = append(message, "And all for the want of a "+rhyme[0]+".")
		} else {
			message = append(message, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
		}
	}
	return message
}
