package helpers

import (
	"os"
	"strconv"
)

func Lock() {
	if lockExists() {
		return
	}
	createLock()
}

func IsLock() bool {
	if lockExists() {
		return true
	}
	return false
}

func UpdateLock() {
	if IsLock() {
		CheckErr("error while updating lock", os.WriteFile(
			"LOCK", []byte(strconv.Itoa(os.Getpid())), 0644,
		))
	}
}

func UnLock() {
	removeLock()
}

func ReadLock() int {
	if IsLock() {
		data, err := os.ReadFile("LOCK")
		PureCheckErr(err)
		process, err := strconv.Atoi(string(data))
		PureCheckErr(err)
		return process
	}
	RaiseStdErr("Error in lock!")
	return -1
}

func createLock() {
	file, err := os.Create("LOCK")
	defer file.Close()
	CheckErr("Error while creating lock", err)
}

func lockExists() bool {
	_, err := os.Stat("LOCK")
	if os.IsNotExist(err) {
		return false
	}
	//CheckErr("Error while checking LOCK", err)
	return true
}

func removeLock() {
	if lockExists() {
		CheckErr("Error while removing LOCK file", os.Remove("LOCK"))
	}
}
