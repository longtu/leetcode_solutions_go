package main

type N014LongestPrefix struct {
}

func (this *N014LongestPrefix) prefixOfTwoStrings(prefix *[]byte, word string) {
	size1 := len(*prefix)
	size2 := len(word)
	endIndex := 0
	for endIndex < size1 && endIndex < size2 && (*prefix)[endIndex] == word[endIndex] {
		endIndex++
	}

	*prefix = (*prefix)[0:endIndex]
}

func (this *N014LongestPrefix) longestCommonPrefix(strs []string) string {
	prefix := make([]byte, 0)
	size := len(strs)
	if size == 0 {
		return string(prefix)
	}

	prefix = []byte(strs[0])
	if size == 1 {
		return string(prefix)
	}

	for i := 1; i < size; i++ {
		this.prefixOfTwoStrings(&prefix, strs[i])
	}

	return string(prefix)
}
