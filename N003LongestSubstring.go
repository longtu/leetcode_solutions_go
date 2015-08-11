package main

type N003LongestSubstring struct {
	stat map[byte]int
}

func (this *N003LongestSubstring) resetStatForStringIndexRange(s string, start int, end int) {
	size := len(s)
	for start < size && end < size && start <= end {
		delete(this.stat, s[start])
		start++
	}

}

func (this *N003LongestSubstring) lengthOfLongestSubstring(s string) int {
	this.stat = make(map[byte]int)
	maxLength := 0
	curLength := 0
	size := len(s)
	start := 0
	end := 0
	for start < size && end < size {
		value, ok := this.stat[s[end]]
		if !ok {
			this.stat[s[end]] = end
			curLength++
			if curLength > maxLength {
				maxLength = curLength
			}
			end++
		} else {
			repeated := value
			this.resetStatForStringIndexRange(s, start, repeated-1)
			this.stat[s[end]] = end
			curLength -= repeated - start
			start = repeated + 1
			end++
		}
	}
	return maxLength
}
