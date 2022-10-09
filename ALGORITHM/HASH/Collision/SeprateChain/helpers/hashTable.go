package helpers

import "fmt"

type HashConfig struct {
	HashSize  int
	hashTable []interface{}
	init      bool
}

func (h *HashConfig) Add(base HashBase) {
	if !h.init {
		h.hashTable = make([]interface{}, h.HashSize)
		h.init = true
	}
	index := int(base.hash(h.HashSize))
	value := base.getData()
	if !h.predictCollision(index) {
		h.hashTable[index] = base.getData()
		return
	}
	h.handleCollision(index, value)
}

func (h *HashConfig) View() {
	for _, ele := range h.hashTable {
		switch ele.(type) {
		case string:
			fmt.Print(ele)
		case LinkedList:
			ls, _ := ele.(LinkedList)
			ls.Show()
		}
	}
}

func (h *HashConfig) predictCollision(index int) bool {
	hashValue := h.hashTable[index]
	if hashValue == nil {
		return false
	}
	return true
}

func (h *HashConfig) handleCollision(index int, data interface{}) {
	hashValue := h.hashTable[index]
	switch hashValue.(type) {
	case string:
		h.AddLinkedList(index, data)
	case LinkedList:
		h.AppendLinkedList(index, data)
	}
}

func (h *HashConfig) AddLinkedList(index int, value interface{}) {
	linkedList := LinkedList{}
	linkedList.Add(value)
	h.hashTable[index] = linkedList
}

func (h *HashConfig) AppendLinkedList(index int, value interface{}) {
	hs := h.hashTable[index]
	v, _ := hs.(LinkedList)
	v.Add(value)
}
