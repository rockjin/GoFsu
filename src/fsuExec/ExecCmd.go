package fsuExec

import (
	"os/exec"
)

func ExecCmd(shellCode string) (out []byte, err error) {
	sh := createSh(shellCode)
	out, err = exec.Command(sh).Output()
	return
}
