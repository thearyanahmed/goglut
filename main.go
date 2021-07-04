package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

const (
	containerPath = "/var/run/glut/containers"
)

func main()  {

	containers := make(map[string]bool)

	fmt.Printf("[+] executing => \n%v\n",os.Args[2:])

	switch os.Args[1] {
	case "run":
		containerID := createContainerID()

		fmt.Printf("container %v\n",containerID)

		containers[containerID] = true
		run(containerID)
	default:
		fmt.Printf("end of the line. am default\n")
	}
}

func run (conID string) {
	cmd := exec.Command(os.Args[2],os.Args[3:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	cmd.Stdin  = os.Stdin

	createDirectoriesIfNotExists(conID)

	cmd.SysProcAttr = &unix.SysProcAttr{
		Chroot: getContanerMountPath(conID),
	}
	
	must(cmd.Run())
}

func createContainerID() string {
	randBytes := make([]byte, 6)
	rand.Read(randBytes)
	return fmt.Sprintf("%02x%02x%02x%02x%02x%02x",
		randBytes[0], randBytes[1], randBytes[2],
		randBytes[3], randBytes[4], randBytes[5])
}

func createDirectoriesIfNotExists(conID string)  {
	homeDir := containerPath + "/" + conID

	contDirs := []string{homeDir + "/fs", homeDir + "/fs/mnt", homeDir + "/fs/upperdir", homeDir + "/fs/workdir"}

	if err := createDirsIfDontExist(contDirs); err != nil {
		log.Fatalf("Unable to create required directories: %v\n", err)
	}
}

func getContanerMountPath(conID string) string {
	return containerPath + "/" + conID + "/fs/mnt"
}

func createDirsIfDontExist(dirs []string) error {
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				log.Printf("Error creating directory: %v\n", err)
				return err
			}
		}
	}
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatalf("error running cmd %v\n",err.Error())
	}
}