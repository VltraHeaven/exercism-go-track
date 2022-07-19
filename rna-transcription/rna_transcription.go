package strand

func ToRNA(dna string) (rna string) {
	rnaConv := map[rune]rune{
		'A': 'U',
		'C': 'G',
		'G': 'C',
		'T': 'A',
	}
	for _, v := range []rune(dna) {
		rna += string(rnaConv[v])
	}
	return
}
