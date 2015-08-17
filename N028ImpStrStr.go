package main

type N028ImpStrStr struct {
}

func (this *N028ImpStrStr) strStr(haystack string, needle string) int {
	length0 := len(haystack)
	length1 := len(needle)
	if length1 == 0 {
		return 0
	}

	i := 0
	j := 0
	k := 0
	for i < length0 {
		if length1 > length0-i {
			return -1
		}
		j = i
		k = 0
		for j < length0 && k < length1 && haystack[j] == needle[k] {
			j++
			k++
		}
		if k == length1 {
			return i
		}
		i++
	}
	return -1
}
