package main

type N030SubstrWithConcat struct {
}

func (this *N030SubstrWithConcat) findSubstring(s string, words []string) []int {
	results := []int{}
	wordsCount := len(words)
	if wordsCount == 0 {
		return results
	}
	wordLength := len(words[0])
	totalLength := wordsCount * wordLength
	sLength := len(s)

	// Build a map<string, count>
	mymap := make(map[string]int)
	for _, word := range words {
		_, ok := mymap[word]
		if !ok {
			mymap[word] = 1
		} else {
			mymap[word] += 1
		}
	}

	sslice := []byte(s)
	// For each substr, check if it exists in the map
	for i := 0; i <= sLength-totalLength; i++ {
		substr := string(sslice[i : i+wordLength])
		_, ok := mymap[substr]
		if ok {
			check := make(map[string]int)
			//copy(check, mymap)
			for k := range mymap {
				check[k] = mymap[k]
			}
			check[substr] -= 1
			found := true
			for j := 1; j < wordsCount; j++ {
				nextSubstr := string(sslice[i+j*wordLength : i+j*wordLength+wordLength])
				v2, ok2 := check[nextSubstr]
				if !ok2 || v2 <= 0 {
					found = false
					break
				} else {
					check[nextSubstr] -= 1
				}
			}

			if found {
				results = append(results, i)
			}
		}
	}
	return results
}
