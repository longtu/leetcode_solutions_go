package main

import "sort"

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortBytes) Len() int {
	return len(s)
}

type N049Anagrams struct {
}

func (this *N049Anagrams) anagrams(strSlice []string) []string {
	strs := make([][]byte, len(strSlice))
	for i := range strSlice {
		strs[i] = []byte(strSlice[i])
	}
	results := make([]string, 0)
	// Sort each string and add to a map. The key is the sorted string, the value is the index of the first string in its group.
	mymap := make(map[string]int)
	size := len(strs)
	for i := 0; i < size; i++ {
		str := make([]byte, len(strs[i]))
		copy(str, strs[i])
		sort.Sort(sortBytes(str))
		strString := string(str)
		v, ok := mymap[strString]
		if !ok {
			mymap[strString] = i
		} else {
			if v >= 0 {
				results = append(results, string(strs[v]))
				mymap[strString] = -1 // Avoid adding it more than once
			}
			results = append(results, string(strs[i]))
		}
	}
	return results
}
