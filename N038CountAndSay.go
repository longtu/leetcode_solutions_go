package main

import "fmt"

const BUFFER_SIZE = 100000

type N038CountAndSay struct {
}

func (this *N038CountAndSay) countAndFill(src []byte, des []byte) {
	start := 0
	srcCur := 0
	desCur := 0
	for src[srcCur] != 0 {
		start = srcCur
		// Count
		for src[start] == src[srcCur] {
			srcCur++
		}
		tmpStr := fmt.Sprintf("%d", srcCur-start)
		length := len(tmpStr)
		for i := 0; i < length; i++ {
			des[desCur] = tmpStr[i]
			desCur++
		}
		// Say
		des[desCur] = src[start]
		desCur++
	}
	for i := 0; i < BUFFER_SIZE; i++ {
		src[i] = 0
	}
}

func (this *N038CountAndSay) countAndSay(n int) string {
	str1 := make([]byte, BUFFER_SIZE)
	str2 := make([]byte, BUFFER_SIZE)

	if n == 0 {
		return ""
	} else if n == 1 {
		str1[0] = '1'
		return string(str1)
	}
	str1[0] = '1'
	i := 1
	for ; i < n; i++ {
		if i%2 == 1 {
			this.countAndFill(str1, str2)
		} else {
			this.countAndFill(str2, str1)
		}
	}

	var src []byte
	if i%2 == 1 {
		src = str1
	} else {
		src = str2
	}
	result := string(src)
	return result
}
