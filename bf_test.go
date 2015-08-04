package bf

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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

func BenchmarkTen(b *testing.B) {
	m := 10

	//Test two hash function first.
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

	//Start benchmark
	b.ResetTimer()
	retList := CreateBloomFilter(inputList, hashList, m)
	if len(retList) != m {
		fmt.Printf("size error %d, exprect %d", len(retList), m)
	}

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(m)
	if MembershipTest(retList[randIndex], retList, hashList) == false {
		fmt.Printf("The %dth member test failed\n", randIndex)
	}
}

func TestBasicCBF(t *testing.T) {
	cbf := NewCountingBloomFilter(100, 0.01)
	cbf.Add([]byte("foo"))
	cbf.Add([]byte("john"))
	cbf.Add([]byte("tom"))

	if cbf.Test([]byte("tom")) == false {
		t.Errorf("CBF: test failed\n")
	}

	cbf.Remove([]byte("john"))
	if cbf.Test([]byte("john")) == true {
		t.Errorf("CBF: remove failed\n")
	}
}
