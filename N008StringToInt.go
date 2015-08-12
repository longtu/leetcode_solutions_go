package main

type N008StringToInt struct {
}

func (this *N008StringToInt) myAtoi(str string) int32 {
	length := len(str)
	if length == 0 {
		return 0
	}

	var value int64 = 0
	index := 0
	indexE := -1
	isNeg := false
	foundE := false
	for str[index] == ' ' && index < length {
		index++
	}
	if str[index] == '-' {
		isNeg = true
		index++
	} else if str[index] == '+' {
		index++
	}
	for ; index < length; index++ {
		if str[index] == 'e' || str[index] == 'E' {
			foundE = true
			indexE = index
			break
		} else if str[index] >= '0' && str[index] <= '9' {
			value = value*10 + int64(str[index]) - int64('0')
			if !isNeg {
				if value > INT_MAX {
					return INT_MAX
				}
			} else {
				if -value < INT_MIN {
					return INT_MIN
				}
			}
		} else {
			break
		}
	}

	var valueAfterE int64 = 0
	if foundE {
		for indexE = indexE + 1; indexE < length; indexE++ {
			valueAfterE = valueAfterE*10 + int64(str[indexE]) - int64('0')
		}
	}

	oldValue := value
	var i int64 = 0
	for ; i < valueAfterE; i++ {
		value *= 10
		if !isNeg {
			if value > INT_MAX {
				value = oldValue
				break
			}
		} else {
			if -value < INT_MIN {
				value = oldValue
				break
			}
		}
	}
	if isNeg {
		value = -value
	}
	if value > INT_MAX {
		value = INT_MAX
	}
	if value < INT_MIN {
		value = INT_MIN
	}
	return int32(value)
}
