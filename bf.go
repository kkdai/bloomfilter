package bf

//CreateBloomFilter
func CreateBloomFilter(setA []int, hashFns [](func(index int) int), m int) []int {
	filter := make([]int, m) //Init filter with size 'm', contains all zero.
	for _, a := range setA {
		for _, hashFn := range hashFns {
			filterIndex := hashFn(a)
			filter[filterIndex] = 1
		}
	}
	return filter
}

func MembershipTest(element int, filterSet []int, hashFns [](func(index int) int)) bool {
	for _, hasFn := range hashFns {
		filterIndex := hasFn(element)
		if filterSet[filterIndex] != 1 {
			return false
		}
	}

	return true
}
