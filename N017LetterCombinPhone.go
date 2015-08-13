package main

type N017LetterCombinPhone struct {
	strings []string
}

func (this *N017LetterCombinPhone) addDigit(digit byte, results *[][]byte) {
	originalSize := len(*results)
	chars := this.strings[digit-'0']
	charSize := len(chars)
	if originalSize == 0 {
		for i := 0; i < charSize; i++ {
			str := []byte{}
			str = append(str, chars[i])
			*results = append(*results, str)
		}
		return
	}

	temp := make([][]byte, len(*results))
	copy(temp, *results)
	for i, _ := range temp {
		temp[i] = make([]byte, len((*results)[i]))
		copy(temp[i], (*results)[i])
	}
	for i := 0; i < charSize-1; i++ {
		for _, v := range temp {
			*results = append(*results, v)
		}
	}

	for i := 0; i < charSize; i++ {
		for j := 0; j < originalSize; j++ {
			(*results)[i*originalSize+j] = append((*results)[i*originalSize+j], chars[i])
		}
	}
}

func (this *N017LetterCombinPhone) letterCombinations(digits string) []string {
	this.strings = []string{"_", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	results := [][]byte{}
	length := len(digits)
	for i := 0; i < length; i++ {
		this.addDigit(digits[i], &results)
	}
	strs := make([]string, len(results))
	for i, _ := range results {
		strs[i] = string(results[i])
	}
	return strs
}
