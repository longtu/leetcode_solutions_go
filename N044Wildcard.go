package main

type N044Wildcard struct {
}

func (this *N044Wildcard) findSubstrWithQuestion(s string, substr string, index int) int {
	//cout << "findSubstrWithQuestion: " << substr << ' ' << index;
	i := index
	length := len(s)
	lenSub := len(substr)
	for ; i < length-lenSub+1; i++ {
		found := true
		for j := 0; j < lenSub; j++ {
			if substr[j] != '?' && s[i+j] != substr[j] {
				found = false
				break
			}
		}
		if found {
			//cout << "found" << endl;
			return i
		}
	}
	//cout << "Not found" << endl;
	return -1
}

func (this *N044Wildcard) hasAllTokens(s string, p []byte, startS int, startP int) bool {
	// Check if p's substrings divided by * exist in s
	i := startS
	j := startP
	lenP := len(p)
	start := j
	end := j
	indexOfPreviousToken := i
	for start < lenP && end < lenP {
		for start < lenP && (p[start] == '*') {
			start++
		}
		if start < lenP {
			end = start
			for end < lenP && p[end] != '*' {
				end++
			}
			if end <= lenP && end > start {
				// Found a token [start, end - 1]
				// Check if this token exists in s
				substr := string([]byte(p)[start:end])
				indexOfPreviousToken = this.findSubstrWithQuestion(s, substr, indexOfPreviousToken)
				if indexOfPreviousToken == -1 {
					return false
				}
				indexOfPreviousToken += end - start
				start = end + 1
				end = end + 1
			}
		}
	}
	return true
}

func (this *N044Wildcard) isMatchInternal(s string, p []byte, startS int, startP int) bool {
	//cout << "Checking " << startS << ' ' << startP << endl;
	i := startS
	j := startP
	lenS := len(s)
	lenP := len(p)
	if lenS-i == 0 && lenP-j == 0 {
		return true
	}
	n := 0
	for k := startP; k < lenP; k++ {
		if p[k] == '*' {
			n++
		}
	}
	if lenS-i < lenP-n-j {
		return false
	}
	for ; j < lenP; j++ {
		if p[j] != '*' {
			if p[j] != '?' {
				if p[j] != s[i] {
					return false // not match
				}
			}
		} else {
			if !this.hasAllTokens(s, p, i, j) {
				return false
			}
			// Expand *
			if this.isMatchInternal(s, p, i, j+1) {
				return true
			}
			for k := 0; k <= lenS-i-(lenP-j-n); k++ {
				if this.isMatchInternal(s, p, i+k+1, j+1) {
					return true
				}
			}
		}
		i++
	}
	if i == lenS && j == lenP {
		return true
	}
	return false
}

func (this *N044Wildcard) find(p []byte, c byte, startIndex int) int {
	index := -1
	for i, size := startIndex, len(p); i < size; i++ {
		if p[i] == c {
			index = i
			break
		}
	}
	return index
}
func (this *N044Wildcard) isMatch(s string, pstr string) bool {
	p := []byte(pstr)
	index := this.find(p, '*', 0)
	if index != -1 {
		lenS := len(s)
		lenP := len(p)
		// Remove adjacent * in p
		for index < lenP && index != -1 {
			for index+1 < lenP && p[index+1] == '*' {
				//p.erase(index, 1);
				temp := p[:index]
				temp = append(temp, p[index+1:]...)
				p = temp
				lenP--
			}
			index = this.find(p, '*', index+1)
		}
		// Scan from end to find the unmatched char
		for i := 0; i < lenP; i++ {
			c := p[lenP-i-1]
			if c == '*' {
				break
			}
			if i == lenS {
				return false
			}
			if c != '?' && c != s[lenS-i-1] {
				return false
			}
		}
	}
	// Scan from begin
	return this.isMatchInternal(s, p, 0, 0)
}
