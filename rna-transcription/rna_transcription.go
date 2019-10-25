package strand

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			rna[i] = 'U'
		case 'C':
			rna[i] = 'G'
		case 'G':
			rna[i] = 'C'
		case 'T':
			rna[i] = 'A'
		}
	}
	return string(rna)
}
