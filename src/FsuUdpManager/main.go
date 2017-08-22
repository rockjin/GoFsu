package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	mdbctx := md5.New()
	mdbctx.Write([]byte("test mdb encrypto"))
	cipherStr := mdbctx.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Println(string(cipherStr))
	fmt.Println(hex.EncodeToString(cipherStr))
}
