package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Println("Nope..")
	}
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	check(cmd.Run())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
