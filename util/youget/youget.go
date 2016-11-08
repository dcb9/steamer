package youget

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func Excecute(arg ...string) (bytes.Buffer, bytes.Buffer) {
	cmd := Cmd(arg...)

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	fmt.Printf("Execute command %q\n", cmd)
	err := cmd.Run()

	fmt.Println("stdout: \n", stdOut.String(), "\n stderr: \n", stdErr.String(), "\n")
	if err != nil {
		log.Fatal("Execute command error: ", err, "\n", stdErr.String())
	}

	return stdOut, stdErr
}

func Cmd(arg ...string) *exec.Cmd {
	cmd := exec.Command("you-get", arg...)
	return cmd
}
