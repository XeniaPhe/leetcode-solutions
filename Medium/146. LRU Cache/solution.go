package lru_cache

type lruNode struct {
    key, value int
    left, right *lruNode
}

type LRUCache struct {
    nodes []lruNode
    dict map[int]*lruNode
    lru, mru *lruNode
}

func Constructor(capacity int) LRUCache {
    return LRUCache{make([]lruNode, capacity), make(map[int]*lruNode, capacity), nil, nil}
}

func (this *LRUCache) use(node *lruNode) {
    if this.mru != node {
        if node.right.left = node.left; this.lru != node {
            node.left.right = node.right
        } else {
            this.lru, this.lru.right.left = this.lru.right, nil
        }

        this.mru, this.mru.right, node.left, node.right = node, node, this.mru, nil
    }
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.dict[key]; exists {
        this.use(node)
        return node.value
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.dict[key]; exists {
        node.value = value
        this.use(node)
    } else if len(this.dict) == 0 {
        node, this.dict[key] = &this.nodes[0], &this.nodes[0]  
        this.lru, this.mru, node.key, node.value = node, node, key, value
    } else if len(this.dict) == len(this.nodes) {
        delete(this.dict, this.lru.key)
        node, this.dict[key] = this.lru, this.lru

        if node.key, node.value = key, value; len(this.dict) != 1 {
            this.lru, this.lru.right.left = this.lru.right, nil
            this.mru, this.mru.right, node.left, node.right = node, node, this.mru, nil
        }
    } else {
        node, this.dict[key] = &this.nodes[len(this.dict)], &this.nodes[len(this.dict)]
        this.mru, this.mru.right, node.left, node.key, node.value = node, node, this.mru, key, value
    }
}