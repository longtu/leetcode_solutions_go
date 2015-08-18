package main

type Number struct {
	m_digits     []byte
	m_isPositive bool
	m_dotIndex   int // From the end. E.g. dotIndex of 2.5 is -1. If digits are 25 and dotIndex is 1, the actual number is 250
}

// vector multiplies a single digit
func (this *Number) multiplyDigit(vec *[]byte, digit byte) {
	size := len(*vec)
	var multi, adv byte
	for i := size - 1; i >= 0; i-- {
		multi = (*vec)[i]*digit + adv
		(*vec)[i] = multi % 10
		adv = multi / 10
	}
	if adv > 0 {
		tmp := make([]byte, size+1)
		tmp[0] = byte(adv)
		for i, size := 0, len(*vec); i < size; i++ {
			tmp[1+i] = (*vec)[i]
		}
		*vec = tmp
	}
}

// vector adds another vector
// vec1 is longer than vec2
func (this *Number) addDigits(pvec1 *[]byte, vec2 []byte) {
	vec1 := *pvec1
	size1 := len(vec1)
	size2 := len(vec2)

	i := 0
	var adv byte
	for ; i < size2; i++ {
		c1 := vec1[size1-i-1] + vec2[size2-i-1] + adv
		vec1[size1-i-1] = c1 % 10
		adv = c1 / 10
	}
	for ; i < size1; i++ {
		c1 := vec1[size1-i-1] + adv
		vec1[size1-i-1] = c1 % 10
		adv = c1 / 10
	}
	if adv > 0 {
		tmp := make([]byte, len(vec1)+1)
		tmp[0] = adv
		for i, size := 0, len(vec1); i < size; i++ {
			tmp[1+i] = vec1[i]
		}
		*pvec1 = tmp
	}
}

func (this *Number) Init(num string) {
	this.m_isPositive = true
	this.m_dotIndex = 0
	length := len(num)
	if length == 0 {
		return
	}
	tempDot := -1
	var c byte
	for i := 0; i < length; i++ {
		c = num[i]
		if c == '+' {
			this.m_isPositive = true
		} else if c == '-' {
			this.m_isPositive = false
		} else if c == '.' {
			tempDot = i
		} else {
			this.m_digits = append(this.m_digits, c-'0')
		}
	}
	if tempDot >= 0 {
		this.m_dotIndex = tempDot + 1 - length
	}

	// Remove leading 0
	leading0Count := 0
	for i := range this.m_digits {
		if this.m_digits[i] == 0 {
			leading0Count++
		} else {
			break
		}
	}
	this.m_digits = this.m_digits[leading0Count:]
}

func (this *Number) multiply(pother *Number) {
	other := *pother
	if (this.m_isPositive == true && other.m_isPositive == true) ||
		(this.m_isPositive == false && other.m_isPositive == false) {
		this.m_isPositive = true
	} else {
		this.m_isPositive = false
	}

	lengthOther := len(other.m_digits)
	if lengthOther == 0 || len(this.m_digits) == 0 {
		this.m_digits = this.m_digits[:0]
		this.m_dotIndex = 0
		return
	}
	copydigits := make([]byte, len(this.m_digits))
	copy(copydigits, this.m_digits)
	for i := 0; i < lengthOther; i++ {
		if i == 0 {
			this.multiplyDigit(&this.m_digits, other.m_digits[i])
		} else {
			this.m_digits = append(this.m_digits, 0)
			temp := make([]byte, len(copydigits))
			copy(temp, copydigits)
			this.multiplyDigit(&temp, other.m_digits[i])
			this.addDigits(&this.m_digits, temp)
		}
	}
	this.m_dotIndex += other.m_dotIndex
}

func (this *Number) toString() string {
	str := make([]byte, 0)
	// Sign
	if !this.m_isPositive {
		str = append(str, '-')
	}

	// Digits and dot
	length := len(this.m_digits)
	if length == 0 {
		str = append(str, '0')
		return string(str)
	}

	if length == -this.m_dotIndex {
		str = append(str, '0')
	}
	for i := 0; i < length; i++ {
		if i == length+this.m_dotIndex {
			// if the remaining digits are all 0, need not add dot
			allZero := true
			for j := i; j < length; j++ {
				if this.m_digits[j] != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				break
			}
			str = append(str, '.')
		}
		str = append(str, this.m_digits[i]+'0')
	}
	// Append possible '0's
	for i := 0; i < this.m_dotIndex; i++ {
		str = append(str, '0')
	}
	return string(str)
}

type N043MultiplyStr struct {
}

func (this *N043MultiplyStr) multiply(num1 string, num2 string) string {
	var n1, n2 Number
	n1.Init(num1)
	n2.Init(num2)
	n1.multiply(&n2)
	return n1.toString()
}
