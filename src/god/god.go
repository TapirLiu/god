package main

import (
   "os"
   "os/exec"
   "path/filepath"
)

func main () {
   gopath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
   os.Setenv ("GOPATH", gopath)
   cmd := exec.Command("go", os.Args [1:]...)
   cmd.Stdin = os.Stdin
	 cmd.Stdout = os.Stdout
	 cmd.Stderr = os.Stderr
   cmd.Run ()
}