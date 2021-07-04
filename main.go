package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main()  {
	fmt.Printf("[+] executing => \n%v\n",os.Args[2:])

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		fmt.Printf("end of the line. am default\n")
	}
}

func run () {
	cmd := exec.Command("/proc/self/exe",append([]string{"child"},os.Args[2:]...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	cmd.Stdin  = os.Stdin

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child () {
	fmt.Printf("running %v as pid %v \n",os.Args[2],os.Getpid())

	cmd := exec.Command(os.Args[2],os.Args[3:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	cmd.Stdin = os.Stdin

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		log.Fatalf("error running cmd %v\n",err.Error())
	}
}
