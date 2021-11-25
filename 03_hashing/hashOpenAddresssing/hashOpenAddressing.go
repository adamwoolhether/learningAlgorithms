package main

import (
	"fmt"
	"hash/fnv"
	"log"
)

type Bucket struct {
	key   string
	value int
}

type Hashtable struct {
	table map[int]Bucket
	size  int
	entry int
}

func NewHashtable(numObjects int) *Hashtable {
	return &Hashtable{

		table: make(map[int]Bucket, numObjects),
		size:  numObjects,
		entry: 0,
	}
}

func (h *Hashtable) get(k string) int {
	hashCode := hash(k) % h.size
	for _, bucket := range h.table {
		if bucket.key == k {
			return bucket.value
		}
		hashCode = (hashCode + 1) % h.size
	}

	return 0
}

func (h *Hashtable) put(k string, v int) {
	hashCode := hash(k) % h.size

	for _, bucket := range h.table {
		if bucket.key == k {
			bucket.value = v
			return
		}
		hashCode = (hashCode + 1) % h.size
	}

	if h.entry >= h.size-1 {
		log.Fatalln("table is full")
	}

	h.table[hashCode] = Bucket{key: k, value: v}
	h.entry++
}

func hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	table := NewHashtable(1000)

	table.put("April", 30)
	table.put("May", 31)
	table.put("September", 30)

	fmt.Println(table.get("August"))
	fmt.Println(table.get("September"))
}