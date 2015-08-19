package main

type N047PermutationsII struct {
}

func (this *N047PermutationsII) permuteUnique(nums []int) [][]int {
	mymap := make(map[int]int) // number, count
	for i := range nums {
		_, ok := mymap[nums[i]]
		if !ok {
			mymap[nums[i]] = 1
		} else {
			mymap[nums[i]] += 1
		}
	}
	results := make([][]int, 0)
	myqueue := make([][]int, 0)
	for k := range mymap {
		newvector := make([]int, 0, 1)
		newvector = append(newvector, k)
		myqueue = append(myqueue, newvector)
	}
	size := len(nums)
	for true {
		candi := myqueue[0]
		if len(candi) == size {
			break
		}
		copymap := make(map[int]int)
		for k := range mymap {
			copymap[k] = mymap[k]
		}
		for i := range candi {
			copymap[candi[i]] -= 1
		}

		// Add number only when the remaining count in the map > 0
		for k := range copymap {
			if copymap[k] > 0 {
				copycandi := make([]int, len(candi), len(candi)+1)
				copy(copycandi, candi)
				copycandi = append(copycandi, k)
				myqueue = append(myqueue, copycandi)
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
