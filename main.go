package main

import (
	"container/list"
	"fmt"
)

var BucketSize int = 89 // prime number

// used for a deterministic match
type Data struct {
	Key   int
	Value int
}

// used for a fast lookup
type MyHashMap struct {
	Bucket []*list.List
}

func Constructor() MyHashMap {
	b := make([]*list.List, BucketSize)
	return MyHashMap{
		Bucket: b,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	// new
	i := key % BucketSize

	if this.Bucket[i] == nil {
		ll := list.New()
		element := Data{Key: key, Value: value}
		ll.PushFront(&element)
		this.Bucket[i] = ll
	} else {
		for e := this.Bucket[i].Front(); e != nil; e = e.Next() {
			ele := (*e).Value.(*Data)
			fmt.Println(ele)
			if ele.Key == key {
				ele.Value = value
				return
			}
		}

		// key is different, so we have a hash collision
		element := Data{Key: key, Value: value}
		this.Bucket[i].PushFront(&element)
	}
}

func (this *MyHashMap) Get(key int) int {
	i := key % BucketSize
	bucket := this.Bucket[i]

	if bucket != nil {

		for e := bucket.Front(); e != nil; e = e.Next() {
			if (*e).Value.(*Data).Key == key {
				return (*e).Value.(*Data).Value
			}
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	i := key % BucketSize
	bucket := this.Bucket[i]
	if bucket != nil {
		for e := bucket.Front(); e != nil; e = e.Next() {
			if (*e).Value.(*Data).Key == key {
				bucket.Remove(e)
			}
		}
	}
}
