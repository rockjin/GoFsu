package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fsuExec"
	"os"
	"path/filepath"

	"fmt"
)

func main() {
	mdbctx := md5.New()
	cipherStr := mdbctx.Sum([]byte("test mdb encrypto"))
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
	fsuExec.ExecCmd("dir")
	println("built in print.")
	var name *string = flag.String("Name", "input your name", "name = 'your name'")
	var age *int = flag.Int("Age", 9, "age=")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*name, *age)
	fmt.Println(filepath.Clean(os.Args[0]))
}
