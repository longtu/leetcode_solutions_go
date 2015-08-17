package main

type N032LongestValidPare struct {
}

func (this *N032LongestValidPare) longestValidParentheses(s string) int {
	mystack := []int{} // Store the index of unmatched '(' or ')', including -1, index of '(', index of unmatched ')'
	size := len(s)
	maxLength := 0
	length := 0
	mystack = append(mystack, -1) //push
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			mystack = append(mystack, i) //push
		} else {
			// Current char is ')'
			mystack = mystack[:len(mystack)-1] //pop
			// The top index after popping is the index of previous unmatched char
			if len(mystack) > 0 {
				length = i - mystack[len(mystack)-1]
				if length > maxLength {
					maxLength = length
				}
			} else {
				// Push index of unmatched ')'
				mystack = append(mystack, i) //push
			}
		}
	}
	return maxLength
}
