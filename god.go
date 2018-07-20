package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	gopath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	os.Setenv("GOPATH", gopath)
	cmd := exec.Command("go", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
