package helpers

import (
	"hash/fnv"
)

type HashBase interface {
	hash(int) uint32
	getData() string
}

func getSha(data string) uint32 {
	h := fnv.New32()
	h.Write([]byte(data))
	return h.Sum32()
}

func GetHash(data string, hashSize int) uint32 {
	hashed := getSha(data)
	return hashed % uint32(hashSize)
}

type StrHash struct {
	Data string
}

func (s StrHash) hash(hashSize int) uint32 {
	return GetHash(s.Data, hashSize)
}

func (s StrHash) getData() string {
	return s.Data
}
