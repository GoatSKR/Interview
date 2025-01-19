package main

// Node represents a node in the doubly linked list
type Node struct {
	key, value int
	prev, next *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head, tail *Node
}

// LRUCache represents the LRU cache
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	list     *DoublyLinkedList
}

// NewLRUCache creates a new LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		list:     &DoublyLinkedList{},
	}
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key int) int {
	if node, found := c.cache[key]; found {
		c.moveToFront(node)
		return node.value
	}
	return -1
}

// Put adds a value to the cache
func (c *LRUCache) Put(key, value int) {
	if node, found := c.cache[key]; found {
		node.value = value
		c.moveToFront(node)
	} else {
		if len(c.cache) >= c.capacity {
			delete(c.cache, c.list.tail.key)
			c.list.remove(c.list.tail)
		}
		node := &Node{key: key, value: value}
		c.cache[key] = node
		c.list.addToFront(node)
	}
}

// moveToFront moves a node to the front of the list
func (c *LRUCache) moveToFront(node *Node) {
	c.list.remove(node)
	c.list.addToFront(node)
}

// addToFront adds a node to the front of the list
func (l *DoublyLinkedList) addToFront(node *Node) {
	node.next = l.head
	node.prev = nil
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

// remove removes a node from the list
func (l *DoublyLinkedList) remove(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}
