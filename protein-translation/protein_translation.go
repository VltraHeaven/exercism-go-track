package protein

import (
	"errors"
)

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop codon")
)

func FromRNA(rna string) (proteins []string, err error) {
	if len(rna)%3 != 0 {
		err = ErrInvalidBase
		return
	}

	for i := range rna {
		if n := i + 1; n%3 == 0 {
			p, e := FromCodon(rna[n-3 : n])
			if e != nil {
				return
			}
			proteins = append(proteins, p)
		}
	}
	return
}

func FromCodon(codon string) (protein string, err error) {
	proteins := map[string]string{
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
	var ok bool
	protein, ok = proteins[codon]
	switch {
	case !ok:
		err = ErrInvalidBase
	case protein == "":
		err = ErrStop
	}
	return
}
