package bf

// This is a basic concept of Bloom Filter, detail implement will be struct bloomfilter

//CreateBloomFilter the basic function to create bloom filter.
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

//MembershipTest the basic test function from a bloom filter list.
func MembershipTest(element int, filterSet []int, hashFns [](func(index int) int)) bool {
	for _, hasFn := range hashFns {
		filterIndex := hasFn(element)
		if filterSet[filterIndex] != 1 {
			return false
		}
	}

	return true
}
