package kata

import "strings"

func DNAtoRNA(dna string) string {
	result := []string{}
	for _, ch := range dna {
		rna := string(ch)
		if ch == 'T' {
			rna = "U"
		}
		result = append(result, rna)
	}
	return strings.Join(result, "")
}
