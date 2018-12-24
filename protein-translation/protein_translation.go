// Package protein implements the "Protein Translation" exercise.
package protein

import "errors"

var (
	STOP           = errors.New("STOP")
	ErrInvalidBase = errors.New("Invalid codon")
)

var translation = map[string](string){
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
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon returns the protein corresponding to the given codon
func FromCodon(in string) (string, error) {
	protein, ok := translation[in]
	if ok {
		if protein == "STOP" {
			return "", STOP
		}
		return protein, nil

	}
	return "", ErrInvalidBase
}

// FromRNA returns the protein sequence corresponding to the given RNA seqeuence
func FromRNA(in string) (out []string, err error) {
	for i := 0; i < len(in)-1; i += 3 {
		slice := in[i : i+3]
		protein, err := FromCodon(slice)

		if err != nil {
			if err == STOP {
				err = nil
			}
			return out, err
		}
		out = append(out, protein)
	}
	return out, nil
}
