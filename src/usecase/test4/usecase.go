package test4

import "regexp"

func CountSmiles(smiles []string) int {
	reg := regexp.MustCompile(`^[:;][-~]?[)D]$`)
	count := 0

	for _, smile := range smiles {
        if reg.MatchString(smile) {
            count++
        }
    }
	
	return count
}