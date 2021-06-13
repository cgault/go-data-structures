package hashtable

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the has
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key   string
	value interface{}
	next  *bucketNode
}

// Insert will take in a key and add it to the has table array
func (h *HashTable) Insert(key string, value interface{}) {
	index := hash(key)
	h.array[index].insert(key, value)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) (interface{}, bool) {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string, v interface{}) {
	if _, ok := b.search(k); !ok {
		newNode := &bucketNode{key: k, value: v}
		newNode.next = b.head
		b.head = newNode
	}
	// already exists
}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) (interface{}, bool) {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return currentNode.value, true
		}
		currentNode = currentNode.next
	}
	return nil, false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode != nil && previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
