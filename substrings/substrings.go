package substrings

func CountVowelSubstrings(word string) int {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}
	count := 0

	for start := 0; start < len(word); start++ {
		uniqueVowels := make(map[byte]bool)
		for end := start; end < len(word); end++ {
			currentCharacter := word[end]
			if !vowels[currentCharacter] {
				break
			}
			uniqueVowels[currentCharacter] = true
			if len(uniqueVowels) == 5 {
				count++
			}
		}
	}

	return count
}
