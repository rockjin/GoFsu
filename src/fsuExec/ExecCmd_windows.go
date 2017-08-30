package fsuExec

import (
	"fmt"
	"os"
)

func createSh(cmd string) string {
	fileName := "test.bat"
	fout, err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		fmt.Println(fileName, err)
		return fileName
	}
	fout.WriteString(cmd)
	return fileName
}
