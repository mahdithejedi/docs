package helpers

import "os"

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

func UnLock() {
	removeLock()
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
	CheckErr("Error while checking LOCK", err)
	return true
}

func removeLock() {
	if lockExists() {
		CheckErr("Error while removing LOCK file", os.Remove("LOCK"))
	}
}
