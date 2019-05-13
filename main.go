package main

import (
	"fmt"
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
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

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

func child() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	check(syscall.Sethostname([]byte("container")))
	check(syscall.Chroot("/home/liz/ubuntufs"))
	check(os.Chdir("/"))
	check(syscall.Mount("proc", "proc", "proc", 0, ""))
	check(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	check(cmd.Run())

	check(syscall.Unmount("proc", 0))
	check(syscall.Unmount("thing", 0))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
