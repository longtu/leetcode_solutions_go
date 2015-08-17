package main

type N029DivideTwoInt struct {
}

// Dividend and divisor should both >0 or both <0
func (this *N029DivideTwoInt) divideInternal(dividend int, divisor int, remainder *int) int {
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		*remainder = 0
		return 0
	}

	i := 0
	if dividend > 0 && divisor > 0 {
		for dividend >= divisor {
			dividend -= divisor
			i++
		}
	} else {
		for dividend <= divisor {
			dividend -= divisor
			i++
		}
	}
	*remainder = dividend
	return i
}

func (this *N029DivideTwoInt) divide(dividend int, divisor int) int {
	BIG_NUMBER := 10000
	if divisor == 0 {
		return INT_MAX
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == INT_MIN {
			return INT_MAX
		}
		return -dividend
	}
	if dividend == 0 {
		return 0
	}
	if divisor == INT_MIN {
		if dividend == INT_MIN {
			return 1
		}
		return 0
	}

	remainder := 0
	quotient := 0
	flip := 1
	if dividend > 0 {
		if divisor < 0 {
			flip = -1
			divisor = -divisor
		}
		if divisor < BIG_NUMBER+1 {
			bigDivisor := 0
			for i := 0; i < BIG_NUMBER; i++ {
				bigDivisor += divisor
			}
			bigQuotient := this.divideInternal(dividend, bigDivisor, &remainder)
			quotient = this.divideInternal(remainder, divisor, &remainder)
			temp := 0
			for i := 0; i < BIG_NUMBER; i++ {
				temp += bigQuotient
			}
			if flip > 0 {
				return (temp + quotient)
			} else {
				return -(temp + quotient)
			}
		}

		quotient = this.divideInternal(dividend, divisor, &remainder)
		if flip > 0 {
			return quotient
		} else {
			return -quotient
		}
	} else {
		// Make sure bothdividend and divisor are negative
		if divisor > 0 {
			flip = -1
			divisor = -divisor
		}
		if divisor > -BIG_NUMBER-1 {
			bigDivisor := 0
			for i := 0; i < BIG_NUMBER; i++ {
				bigDivisor += divisor
			}
			bigQuotient := this.divideInternal(dividend, bigDivisor, &remainder)
			quotient = this.divideInternal(remainder, divisor, &remainder)
			temp := 0
			for i := 0; i < BIG_NUMBER; i++ {
				temp += bigQuotient
			}
			if flip > 0 {
				return (temp + quotient)
			} else {
				return -(temp + quotient)
			}
		}

		quotient = this.divideInternal(dividend, divisor, &remainder)
		if flip > 0 {
			return quotient
		} else {
			return -quotient
		}
	}
}

//int i = INT_MIN;
//cout << (-INT_MIN == INT_MIN) << endl;
