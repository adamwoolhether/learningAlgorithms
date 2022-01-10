package main

import (
	"testing"
)

func Test_NewHashtable(t *testing.T) {
	lengths := []int{5, 10, 20}

	for _, tc := range lengths {
		t.Run("new hashtable creation", func(*testing.T) {
			ht := NewHashtable(tc)
			if ht.table == nil {
				t.Fatal("hash table not created")
			}

			if len(ht.table) > 1 {
				t.Fatal("initial hash table should have no nodes")
			}

			if ht.num != 0 {
				t.Fatal("initialized hash table should have no node count")
			}

			if ht.max != tc {
				t.Fatalf("initialzed hash table should have a length of %d, got: %d", tc, &ht.max)
			}
		})
	}
}
