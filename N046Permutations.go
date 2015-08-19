package main

type N046Permutations struct {
}

func (this *N046Permutations) find(p []int, c int, startIndex int) int {
	index := -1
	for i, size := startIndex, len(p); i < size; i++ {
		if p[i] == c {
			index = i
			break
		}
	}
	return index
}

func (this *N046Permutations) permute(nums []int) [][]int {
	results := make([][]int, 0)
	myqueue := make([][]int, 0)
	for i := range nums {
		newvector := make([]int, 0, 1)
		newvector = append(newvector, nums[i])
		myqueue = append(myqueue, newvector)
	}
	size := len(nums)
	for true {
		candi := myqueue[0]
		if len(candi) == size {
			break
		}
		for i := range nums {
			if this.find(candi, nums[i], 0) == -1 {
				copycandi := make([]int, len(candi), len(candi)+1)
				copy(copycandi, candi)
				copycandi = append(copycandi, nums[i])
				myqueue = append(myqueue, copycandi) // Push
			}
		}
		myqueue = myqueue[1:] // Pop
	}
	for len(myqueue) > 0 {
		results = append(results, myqueue[0])
		myqueue = myqueue[1:] // Pop
	}
	return results
}
