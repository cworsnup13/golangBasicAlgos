package hashTable

import (
	"code.uber.internal/personal/golangBasicAlgos/src/dataStructures/linkedList"
)

const (
	DefaultHashTableSize = 32
	primeVal = 31
)

type key string

type HashItem struct {
	k key
	v interface{}
}

type HashTable interface {
	Set(key, interface{})
	Delete(key) *linkedList.Item
	Get(key) interface{}
	Has(key) bool
	GetKeys() []key
}

type hashTable struct {
	hashTableSize int
	buckets []linkedList.LinkedList
	keys map[key]int
}

func New(size int) HashTable {
	ll := make([]linkedList.LinkedList, size)
	for i := range ll {
		ll[i] = linkedList.New()
	}
	var keys = make(map[key]int)
	h := hashTable{
		hashTableSize: size,
		buckets: ll,
		keys: keys,
	}

	return &h
}

func (h *hashTable) hash(k key) int {
	// hash = charCodeAt(0) * PRIME^(n-1) + charCodeAt(1) * PRIME^(n-2) + ... + charCodeAt(n-1)
	var hash int
	for i, v := range k {
		hash += int(v) * primeVal ^ (len(k)-(i+1))
	}
	return hash % h.hashTableSize
}

func (h *hashTable) Set(k key, v interface{}){
	keyHash := h.hash(k)
	h.keys[k] = keyHash
	bucketLinkedList := h.buckets[keyHash]

	callback := func(nodeval interface{}) bool {return nodeval.(HashItem).k == k}
	node := bucketLinkedList.Find(nil, callback)

	if node == nil {
		bucketLinkedList.Append(HashItem{k, v})
	} else {
		a := node.Value.(HashItem)
		a.v = v
	}

}

func (h *hashTable) Delete(k key) *linkedList.Item {
	keyHash := h.hash(k)
	delete(h.keys, k)
	bucketLinkedList := h.buckets[keyHash]

	callback := func(nodeval interface{}) bool {return nodeval.(HashItem).k == k}
	node := bucketLinkedList.Find(nil, callback)

	if node != nil {
		return bucketLinkedList.Delete(node.Value)
	}

	return nil
}

func (h *hashTable) Get(k key) interface{} {
	bucketLinkedList := h.buckets[h.hash(k)]
	callback := func(nodeval interface{}) bool {return nodeval.(HashItem).k == k}
	node := bucketLinkedList.Find(nil, callback)
	if node != nil {
		return node.Value.(HashItem).v
	}
	return nil
}

func (h *hashTable) Has(k key) bool {
	for v := range h.keys {
		if v == k {
			return true
		}
	}
	return false
}

func (h *hashTable) GetKeys() []key {
	var keys = make([]key, len(h.keys))
	var i = 0
	for k := range h.keys {
		keys[i] = k
		i++
	}
	return keys
}