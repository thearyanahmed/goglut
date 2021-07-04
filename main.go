package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main()  {
	fmt.Printf("[+] executing => \n%v\n",os.Args[2:])

	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Printf("end of the line. am default\n")
	}
}

func run () {
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