package test2

func Permutations(inputWord string) []string  {
	
	if len(inputWord) == 1 {
		return []string{inputWord}
	}

	newWords := []string{}
	for i, word := range inputWord {
		result := inputWord[:i] + inputWord[i+1:]
		for _, p := range Permutations(result) {
			newWords = append(newWords, string(word) + p )
		}
	}
	
	return newWords
}
