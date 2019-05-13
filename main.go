package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "fork":
		fork()
	default:
		panic("nope")
	}
}

func run() {
	fmt.Printf("run() executing %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"fork"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func fork() {
	fmt.Printf("fork() executing %v \n", os.Args)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("gotain")))
	must(syscall.Chroot(fmt.Sprintf("%s/ubuntufs", os.Getenv("HOME"))))
	must(os.Chdir("/"))

	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())
	must(syscall.Unmount("proc", 0))

}

func must(err error) {
	if err != nil {
		fmt.Printf("has error: %v \n", err)
	}
}
