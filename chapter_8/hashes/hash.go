package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	// create a new hash
	h := crc32.NewIEEE()
	// write data to hash
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

	// use a sha1 hash
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte("test"))
	bs := sha1Hash.Sum([]byte{})
	fmt.Println(bs)

}
