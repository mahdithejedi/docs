package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"
)

var chars = strings.Split("abcdefghijklmonpqrstuvwxyzABCDEFGHIJKLMONPQRSTUVWXYZ", "")

func Hash(url string) string {
	hasher := sha1.New()
	io.WriteString(hasher, url)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))

}

func Shorten(url string, salt string) string {
	hashed_url := Hash(url + salt)
	return hashed_url[:5]

}

func GetSalt() string {
	salt_len := len(chars)
	output := []string{}
	herb := int(time.Now().Unix())
	for ; herb < salt_len; herb = int(herb / salt_len) {
		output = append(output, chars[int(herb%salt_len)])
	}
	return strings.Join(output, "")
}

func main() {
	fmt.Println(GetSalt())
	// fmt.Println(Shorten("google.com", "aa"))
}
