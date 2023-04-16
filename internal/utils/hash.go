package utils

import "hash/fnv"

func Hash(b []byte) (uint32, error) {
	h := fnv.New32a()
	_, err := h.Write(b)
	return h.Sum32(), err
}
