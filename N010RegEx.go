package main

type N010RegEx struct {
}

func (this *N010RegEx) findSubstrWithDot(s string, substr string, index int) int {
	//cout << "findSubstrWithDot: " << substr << ' ' << index;
	i := index
	lenS := len(s)
	lenSub := len(substr)
	for ; i < lenS-lenSub+1; i++ {
		found := true
		for j := 0; j < lenSub; j++ {
			if substr[j] != '.' && s[i+j] != substr[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func (this *N010RegEx) hasAllTokens(s string, p string, startS int, startP int) bool {
	// Check if p's substrings divided by * exist in s
	i := startS
	j := startP
	lenP := len(p)
	start := j
	end := j
	indexOfPreviousToken := i
	for start < lenP && end < lenP {
		for start+1 < lenP && p[start+1] == '*' {
			start += 2
		}
		if start < lenP {
			end = start
			for end+1 < lenP && p[end+1] != '*' {
				end++
			}
			if end+1 == lenP {
				end++
			}
			if end <= lenP && end > start {
				// Found a token [start, end - 1]
				// Check if this token exists in s
				substr := p[start:end]
				indexOfPreviousToken = this.findSubstrWithDot(s, string(substr), indexOfPreviousToken)
				if indexOfPreviousToken == -1 {
					return false
				}
				indexOfPreviousToken += end - start
				start = end + 2
				end = end + 2
			}
		}
	}
	return true
}

func (this *N010RegEx) isMatchInternal(s string, p string, startS int, startP int) bool {
	i := startS
	j := startP
	lenS := len(s)
	lenP := len(p)
	if lenS-i == 0 && lenP-j == 0 {
		return true
	}
	n := 0
	for i := startP; i < lenP; i++ {
		if p[i] == '*' {
			n++
		}
	}
	if lenS-i < lenP-2*n-j {
		return false
	}
	for ; j < lenP; j++ {
		if j+1 < lenP {
			if p[j+1] != '*' {
				if p[j] != '.' && p[j] != s[i] {
					return false // not match
				}
			} else {
				if !this.hasAllTokens(s, p, i, j) {
					return false
				}
				// Expand *
				// 0 char
				if this.isMatchInternal(s, p, i, j+2) {
					return true
				}
				// k chars
				for k := 1; k <= lenS-i-(lenP-j-2*n); k++ {
					matched := true
					for q := 0; q < k; q++ {
						if p[j] != '.' && p[j] != s[i+q] {
							matched = false
							break
						}
					}
					if matched {
						if this.isMatchInternal(s, p, i+k, j+2) {
							return true
						}
					}
				}
				// 0 char ~ k chars not match, return false
				return false
			}
		} else {
			if p[j] != '.' && p[j] != s[i] {
				return false
			}
		}
		i++
	}
	if i == lenS && j == lenP {
		return true
	}
	return false
}

func (this *N010RegEx) isMatch(s string, p string) bool {
	// Scan from begin
	return this.isMatchInternal(s, p, 0, 0)
}
