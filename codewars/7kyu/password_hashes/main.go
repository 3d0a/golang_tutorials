package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func PassHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println(PassHash("hello"))
}
