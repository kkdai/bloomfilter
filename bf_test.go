package bf

import (
	"fmt"
	"testing"
)

func ExampleBasicBloomFilter() {
	m := 5
	h1 := func(input int) int {
		return (input % m) + 1
	}

	var hashList [](func(index int) int)
	hashList = append(hashList, h1)

	var inputList = []int{100, 121, 456, 121, 987}
	fmt.Println("Bloom Filter:", CreateBloomFilter(inputList, hashList, m))
}

func TestBasicFun(t *testing.T) {
	m := 10
	h1 := func(input int) int {
		return (input % m) + 1
	}

	h2 := func(input int) int {
		return (input % m) + 1
	}

	var hashList [](func(index int) int)
	hashList = append(hashList, h1)
	hashList = append(hashList, h2)

	var inputList = []int{100, 150, 121, 422, 456, 121, 987, 111, 121}
	retList := CreateBloomFilter(inputList, hashList, m)
	fmt.Println(retList)
	if len(retList) != m {
		t.Errorf("size error %d, exprect %d", len(retList), m)
	}

	for index, element := range retList {
		if MembershipTest(element, retList, hashList) == false {
			t.Errorf("The %dth member test failed", index)
		}
	}
}
