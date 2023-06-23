package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Syntax: main.go <hash_type> <text>")
		return
	}

	hashType, text := os.Args[1], os.Args[2]

	var hash []byte
	switch hashType {
	case "md5":
		h := md5.Sum([]byte(text))
		hash = h[:]
	case "sha1":
		h := sha1.Sum([]byte(text))
		hash = h[:]
	case "sha256":
		h := sha256.Sum256([]byte(text))
		hash = h[:]
	case "sha512":
		h := sha512.Sum512([]byte(text))
		hash = h[:]
	default:
		fmt.Println("Invalid hash type")
		return
	}

	fmt.Printf("%x\n", hash)
}
