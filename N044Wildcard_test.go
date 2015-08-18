package main

import (
	"fmt"
	"testing"
)

func TestN044Wildcard(t *testing.T) {
	var n044 N044Wildcard
	fmt.Print("N044Wildcard:\t\t")
	fmt.Print(n044.isMatch("mississippi", "m*iss*iss*"), " ")
	fmt.Print(n044.isMatch("aaabababbba", "**a*b*bbb*a**"), " ")
	fmt.Print(n044.isMatch("aaaabaabaabbbabaabaabbbbaabaaabaaabbabbbaaabbbbbbabababbaabbabbbbaababaaabbbababbbaabbbaabbaaabbbaabbbbbaaaabaaabaabbabbbaabababbaabbbabababbaabaaababbbbbabaababbbabbabaaaaaababbbbaabbbbaaababbbbaabbbbb", "**a*b*b**b*b****bb******b***babaab*ba*a*aaa***baa****b***bbbb*bbaa*a***a*a*****a*b*a*a**ba***aa*a**a*"), " ")
	fmt.Print(n044.isMatch("ho", "**ho"), " ")
	fmt.Println(n044.isMatch("", "*"), n044.isMatch("aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b"), n044.isMatch("abefcdgiescdfimde", "ab*cd?i*de"), n044.isMatch("bbbaaabaababbabbbaabababbbabababaabbaababbbabbbabb", "*b**b***baba***aaa*b***"), n044.isMatch("ab", "?*"), n044.isMatch("aab", "c*a*b"))
}
