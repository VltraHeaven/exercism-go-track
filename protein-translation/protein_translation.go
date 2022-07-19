package protein
import (
	"errors"
)
var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop codon")
	aminoAcids     = map[string]string{
    	"AUG": "Methionine",
        "UUU": "Phenylalanine",
        "UUC": "Phenylalanine",
        "UUA": "Leucine",
        "UUG": "Leucine",
        "UCU": "Serine",
        "UCC": "Serine",
        "UCA": "Serine",
        "UCG": "Serine",
        "UAU": "Tyrosine",
        "UAC": "Tyrosine",
        "UGU": "Cysteine",
        "UGC": "Cysteine",
        "UGG": "Tryptophan",
        "UAA": "",
        "UAG": "",
        "UGA": "",
    }
)
func FromRNA(rna string) (aminos []string, err error) {
	if len(rna)%3 != 0 {
		err = ErrInvalidBase
		return
	}
	for i := range rna {
		if n := i + 1; n%3 == 0 {
			a, e := FromCodon(rna[n-3 : n])
			if e != nil {
				return
			}
			aminos = append(aminos, a)
		}
	}
	return
}
func FromCodon(codon string) (amino string, err error) {
	var ok bool
	amino, ok = aminoAcids[codon]
	switch {
	case !ok:
		err = ErrInvalidBase
	case amino == "":
		err = ErrStop
	}
	return
}