package main

type N012IntToRoman struct {
}

func (this *N012IntToRoman) digitToRoman(num int, position int, highChar byte, halfChar byte, thisChar byte, pstr *[]byte) int {
	var nine int
	var five int
	var four int
	var one int
	switch position {
	case 3:
		nine = 900
		five = 500
		four = 400
		one = 100
	case 2:
		nine = 90
		five = 50
		four = 40
		one = 10
	case 1:
		nine = 9
		five = 5
		four = 4
		one = 1
	default:
		return 0
	}

	if num >= nine {
		*pstr = append(*pstr, thisChar, highChar)
		num = num % one
	} else if num >= five {
		*pstr = append(*pstr, halfChar)
		for i := 0; i < (num-five)/one; i++ {
			*pstr = append(*pstr, thisChar)
		}
		num = num % one
	} else if num >= four {
		*pstr = append(*pstr, thisChar, halfChar)
		num = num % one
	} else if num >= one {
		for i := 0; i < num/one; i++ {
			*pstr = append(*pstr, thisChar)
		}
		num = num % one
	}
	return num
}

func (this *N012IntToRoman) intToRoman(num int) string {
	str := make([]byte, 0)
	if num >= 1000 {
		for i := 0; i < num/1000; i++ {
			str = append(str, 'M')
		}
		num = num % 1000
	}

	num = this.digitToRoman(num, 3, 'M', 'D', 'C', &str)
	num = this.digitToRoman(num, 2, 'C', 'L', 'X', &str)
	num = this.digitToRoman(num, 1, 'X', 'V', 'I', &str)
	return string(str)
}
