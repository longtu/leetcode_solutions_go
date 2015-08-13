package main

type N020ValidParentheses struct {
}

func (this *N020ValidParentheses) isValid(s string) bool {
	mystack := make([]byte, 0)
	length := len(s)
	var curChar byte
	var topChar byte
	for i := 0; i < length; i++ {
		curChar = s[i]
		if curChar == '(' || curChar == '[' || curChar == '{' {
			mystack = append(mystack, curChar)
		} else if curChar == ')' || curChar == ']' || curChar == '}' {
			stackLen := len(mystack)
			if stackLen == 0 {
				return false
			}
			topChar = mystack[stackLen-1]
			if (curChar == ')' && topChar == '(') || (curChar == ']' && topChar == '[') || (curChar == '}' && topChar == '{') {
				mystack = mystack[:stackLen-1]
			} else {
				mystack = append(mystack, curChar)
			}
		}
	}

	if len(mystack) == 0 {
		return true
	}
	return false
}
