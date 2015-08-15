package main

type N022GenParentheses struct {
}

func (this *N022GenParentheses) generateParenthesis(n int) []string {
	original := []byte{}
	// Build the original string with n '('
	for i := 0; i < n; i++ {
		original = append(original, '(')
	}
	myset := make(map[string]int)
	uniset := make(map[string]int)
	myqueue := []string{string(original)}
	for len(myqueue) != 0 {
		str := myqueue[0]
		myqueue = myqueue[1:]
		if len(str) == n*2 {
			myset[str] = 1
			continue
		}
		length := len(str)
		sum := 0
		var c byte = 0
		// Try inserting ')' to keep sum >= 0 at correct position
		for i := 0; i < length; i++ {
			c = str[i]
			if c == '(' {
				sum++
			} else {
				sum--
			}
			if sum > 0 && (i == length-1 || str[i+1] != ')') {
				str2 := []byte(str)
				//str2 = str2.insert(i + 1, ")");
				tmp := []byte(str)
				tmp = append(tmp[:i+1], ')')
				str2 = append(tmp, str2[i+1:]...)
				sum2 := sum - 1
				isNeg := false
				// Recalculate the sum to make sure sum >= 0, and skip if part sum < 0
				for j := i + 2; j < length+1; j++ {
					c = str2[j]
					if c == '(' {
						sum2++
					} else {
						sum2--
					}
					if sum2 < 0 {
						isNeg = true
						break
					}
				}
				_, ok := uniset[string(str2)]
				if !isNeg && !ok {
					myqueue = append(myqueue, string(str2))
					uniset[string(str2)] = 1
				}
			}

		}
	}

	result := []string{}
	for k := range myset {
		result = append(result, k)
	}
	return result
}
