package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func crc32Hash(bs []byte) uint32 {
	h := crc32.NewIEEE()
	h.Write(bs)
	v := h.Sum32()
	return v
}

func fileHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return crc32Hash(bs), nil
}

func main() {
	bs := []byte("test")
	v := crc32Hash(bs)
	fmt.Println(v)

	h1, err := fileHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := fileHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
