package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
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
	check(syscall.Chroot("/home/gotain/ubuntufs"))
	check(os.Chdir("/"))
	check(syscall.Mount("proc", "proc", "proc", 0, ""))
	check(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	check(cmd.Run())

	check(syscall.Unmount("proc", 0))
	check(syscall.Unmount("thing", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "gotain"), 0755)
	check(ioutil.WriteFile(filepath.Join(pids, "gotain/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	check(ioutil.WriteFile(filepath.Join(pids, "gotain/notify_on_release"), []byte("1"), 0700))
	check(ioutil.WriteFile(filepath.Join(pids, "gotain/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
