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

// __DOCUMENTATION__, __NAMESPACE__
// what is name space? = 'https://man7.org/linux/man-pages/man7/namespaces.7.html'
// what is user namespace? = 'https://man7.org/linux/man-pages/man7/user_namespaces.7.html'
// what is pid namespace? = 'https://man7.org/linux/man-pages/man7/pid_namespaces.7.html'

// __source__ = 'https://youtu.be/8fi7uSYlOdc'

// Namespace => what can you see?
// Cgroup (Control Group) => what can you do?
// Location /sys/fs/cgroup => different Cgroup

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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID, /* | syscall.CLONE_NEWUSER | syscall.CLONE_NEWIPC */
	}
	checkError(cmd.Run())
	fmt.Println("------------------------------------")
}

func child(command string, args ...string) {
	fmt.Println("Running", os.Args[2:], "as", os.Getpid())
	fmt.Println("------------------------------------")
	setSyscall()
	setCgroup()
	runChild(command, args...)
}

func setSyscall() {
	// What is chroot?
	// https://www.geeksforgeeks.org/chroot-command-in-linux-with-examples/
	syscall.Sethostname([]byte("child"))
	syscall.Chroot("/")
	syscall.Chdir("/home/m.moosavi/")

}

func setCgroup() {
	// Limit number of processes to 20
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "liz"), 0755)
	checkError(ioutil.WriteFile(filepath.Join(pids, "liz/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	checkError(ioutil.WriteFile(filepath.Join(pids, "liz/notify_on_release"), []byte("1"), 0700))
	// add current process into this group
	checkError(ioutil.WriteFile(filepath.Join(pids, "liz/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
	// If you run command on shell
	//	:() { : | : & } ; :
	// Then you system won't get fucked because you have limited processes number up to 20
}

func runChild(command string, args ...string) {
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
