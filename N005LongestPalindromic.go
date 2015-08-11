package main

type N005LongestPalindromic struct {
}

func (this *N005LongestPalindromic) findSeeds(s string, pseeds *[]int) {
	length := len(s)
	for i := 0; i < length-1; i++ {
		// Check 2 chars, then 3 chars
		if s[i] == s[i+1] {
			*pseeds = append(*pseeds, i*1000+i+1)
		}
		if i < length-2 && s[i] == s[i+2] {
			*pseeds = append(*pseeds, i*1000+i+2)
		}
	}
}

// Scan for seeds first, and then expand seeds for the longest one
func (this *N005LongestPalindromic) longestPalindrome(s string) string {
	seeds := []int{}
	length := len(s)
	if length <= 1 || (length == 2 && s[0] == s[1]) {
		return s
	}

	this.findSeeds(s, &seeds)
	if len(seeds) == 0 {
		return s[0:1]
	}

	firstSeed := seeds[0]
	begin := firstSeed / 1000
	end := firstSeed % 1000
	maxLength := end - begin + 1
	maxBegin := begin
	maxEnd := end
	for i, size := 0, len(seeds); i < size; i++ {
		begin = seeds[i] / 1000
		end = seeds[i] % 1000

		for begin > 0 && end < length-1 && s[begin-1] == s[end+1] {
			begin--
			end++
		}

		if end-begin+1 > maxLength {
			maxLength = end - begin + 1
			maxBegin = begin
			maxEnd = end
		}
	}

	return s[maxBegin : maxEnd+1]
}
