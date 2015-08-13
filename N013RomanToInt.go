package main

type N013RomanToInt struct {
}

func (this *N013RomanToInt) romanToInt(s string) int {
	value := 0
	length := len(s)
	var current byte
	var pre byte
	for i := length - 1; i >= 0; {
		current = s[i]
		if i-1 >= 0 {
			pre = s[i-1]
		} else {
			pre = 0
		}
		if current == 'I' {
			value += 1
			i--
		} else if pre == 'I' && current == 'V' {
			value += 4
			i -= 2
		} else if current == 'V' {
			value += 5
			i--
		} else if pre == 'I' && current == 'X' {
			value += 9
			i -= 2
		} else if current == 'X' {
			value += 10
			i--
		} else if pre == 'X' && current == 'L' {
			value += 40
			i -= 2
		} else if current == 'L' {
			value += 50
			i--
		} else if pre == 'X' && current == 'C' {
			value += 90
			i -= 2
		} else if current == 'C' {
			value += 100
			i--
		} else if pre == 'C' && current == 'D' {
			value += 400
			i -= 2
		} else if current == 'D' {
			value += 500
			i--
		} else if pre == 'C' && current == 'M' {
			value += 900
			i -= 2
		} else if current == 'M' {
			value += 1000
			i--
		} else {
			break
		}
	}
	return value
}
