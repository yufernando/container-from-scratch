package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker run <container> cmd args
// go     run main.go run cmd args

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what??")
	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stout = os.Stdout
	cmd.Sterr = os.Stderr

	must(cmd.Run())
}

func mist(err error) {
	if err != nil {
		panic(err)
	}
}