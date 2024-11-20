package protein

import (
	"errors"
)

var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")
var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var res []string
	for i := 0; i <= len(rna)-3; i += 3 {
		aa, ok := FromCodon(rna[i : i+3])
		if ok == ErrInvalidBase {
			return res, ErrInvalidBase
		} else if ok == ErrStop {
			return res, nil
		}

		res = append(res, aa)
	}

	return res, nil
}

func FromCodon(codon string) (string, error) {
	aa, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	} else if aa == "STOP" {
		return "", ErrStop
	} else {
		return aa, nil
	}
}
