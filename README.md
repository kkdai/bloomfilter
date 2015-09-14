Bloom Filter
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/bloomfilter/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/bloomfilter?status.svg)](https://godoc.org/github.com/kkdai/bloomfilter)  [![Build Status](https://travis-ci.org/kkdai/bloomfilter.svg?branch=master)](https://travis-ci.org/kkdai/bloomfilter)


[Bloom Filter](https://en.wikipedia.org/wiki/Bloom_filter) is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970.

This implement is experimental to written Golang to prove the basic concept of Bloom Filter.

Install
---------------
`go get github.com/kkdai/bloomfilter`


Usage
---------------

```go
    //Create a couting bloom filter expect size 100, false detect rate 0.01
	cbf := NewCountingBloomFilter(100, 0.01)
	
	//Add item into cbf
	cbf.Add([]byte("foo"))
	cbf.Add([]byte("john"))
	cbf.Add([]byte("tom"))

	//test 
	fmt.Println("Test cbf:", cbf.Test([]byte("tom"))) //return "true"

    //Remvoe item from cbf
	cbf.Remove([]byte("john"))
	
	//test again
	fmt.Println("Test cbf:", cbf.Test([]byte("john"))) //return "false"
```


Inspired
---------------

- [https://github.com/pmylund/go-bloom](https://github.com/pmylund/go-bloom)
- [https://github.com/willf/bloom](https://github.com/willf/bloom)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

