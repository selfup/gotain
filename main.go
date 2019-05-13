package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		log.Fatal("cmd !supported")
	}
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	// compile target for linux because of `Cloneglags`
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	check(cmd.Run())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
