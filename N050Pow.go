package main

type N050Pow struct {
}

func (this *N050Pow) powPos(x float64, n int) float64 {
	if x == 1 {
		return 1
	}
	temp := float64(0)
	result := x
	for i := 1; i < n; i++ {
		temp = result * x
		if temp == result {
			break
		}
		result = temp
	}
	return result
}

func (this *N050Pow) myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}

	if n > 0 {
		if x > 0 {
			return this.powPos(x, n)
		} else {
			if n%2 == 0 {
				return this.powPos(-x, n)
			} else {
				return -this.powPos(-x, n)
			}
		}
	} else {
		if x > 0 {
			return 1.0 / this.powPos(x, -n)
		} else {
			if n%2 == 0 {
				return 1.0 / this.powPos(-x, -n)
			} else {
				return -1.0 / this.powPos(-x, -n)
			}
		}
	}
}
