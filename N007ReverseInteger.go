package main

type N007ReverseInteger struct {
}

func (this *N007ReverseInteger) reverse(x int32) int32 {
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}

	var y int64 = 0
	for x > 0 {
		y = y*10 + int64(x)%10
		x /= 10
	}

	if y > INT_MAX { // Handle overflow
		return 0
	}
	if isNeg {
		y = -y
	}
	return int32(y)

}
