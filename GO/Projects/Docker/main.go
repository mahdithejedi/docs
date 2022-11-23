package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// __DOCUMENTATION__, __NAMESPACE__
// what is name space? = 'https://man7.org/linux/man-pages/man7/namespaces.7.html'
// what is user namespace? = 'https://man7.org/linux/man-pages/man7/user_namespaces.7.html'
// what is pid namespace? = 'https://man7.org/linux/man-pages/man7/pid_namespaces.7.html'

// __source__ = 'https://youtu.be/8fi7uSYlOdc'

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func run(command string, args ...string) {
	fmt.Println("Running", os.Args[2:], "as", os.Getpid())
	fmt.Println("------------------------------------")
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, append([]string{command}, args...)...)...)
	//cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
		//Unshareflags: syscall.CLONE_NEWPID,
	}
	//syscall.Sethostname([]byte("parent"))
	checkError(cmd.Run())
	fmt.Println("------------------------------------")
}

func child(command string, args ...string) {
	fmt.Println("Running", os.Args[2:], "as", os.Getpid())
	fmt.Println("------------------------------------")
	syscall.Sethostname([]byte("child"))
	syscall.Chroot("/")
	syscall.Chdir("/home/m.moosavi/")
	//syscall.Mount("proc", "proc", "proc", 0, "")
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags: syscall.CLONE_NEWUTS,
	//}

	checkError(cmd.Run())
	syscall.Unmount("proc", 0)
	fmt.Println("------------------------------------")
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("you should input a command")
	}
	switch os.Args[1] {
	case "run":
		run(os.Args[2], os.Args[3:]...)
	case "child":
		child(os.Args[2], os.Args[3:]...)
	default:
		log.Fatal("Bad argument")
	}
}
