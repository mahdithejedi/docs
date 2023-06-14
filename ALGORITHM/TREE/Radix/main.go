package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	isLeaf   bool
	children []*Node
	ip       string
}

func fatalError(err error, crash bool) {
	fmt.Fprintln(os.Stderr, err.Error())
	if crash == true {
		os.Exit(0)
	}
}

func printError(err error) {
	fatalError(err, false)
}

func validateIPV4(ip string) bool {
	subIPS := strings.Split(ip, ".")
	if len(subIPS) != 4 {
		printError(errors.New("IPV4 len should be 4"))
		return true
	}
	for subIP := range subIPS {
		if int(subIP) < 0 || int(subIP) > 255 {
			printError(errors.New("subIP range should be between 0-255"))
			return true
		}
	}
	return false
}

func (n *Node) Add(ip string) {
	if validateIPV4(ip) {
		return
	}

	if n.ip == "" {
		n.Initiate(ip)
		return
	}
	lastNode, isExact, remainingIP := n.findPrefix(ip)
	fmt.Println("for IP ", ip, "Last node is", lastNode, "is exact is", isExact, "remainingIP is", remainingIP)
	if isExact {
		fmt.Println("IS exact ofr ip ", ip, "is ", isExact)
		return
	}
	lastNode.children = append(lastNode.children, &Node{isLeaf: true, ip: remainingIP})

}

func (n *Node) exists(ip string) bool {
	_, isExact, _, _ := n.findPrefix(ip)
	if isExact {
		return true
	}
	return false
}

func longestPrefix(str1, str2 string) string {
	var commonPrefix string
	minLength := len(str1)
	if len(str2) < minLength {
		minLength = len(str2)
	}
	for i := 0; i < minLength; i++ {
		if str1[i] != str2[i] {
			break
		}
		commonPrefix += string(str1[i])
	}
	return commonPrefix
}

func (n *Node) findPrefix(ip string) (latestNode *Node, isExact bool, isPartial bool, sharePartial string, latestIp string) {
	head := n
	if head.ip == ip {
		return head, true, false, "", ""
	}
	for _, child := range head.children {
		if child.isLeaf {
			if child.ip == ip {
				return child, true, false, "", ""
			}
			return child, false, false, ip
		}
		if strings.HasPrefix(child.ip, ip) {
			ip, _ = strings.CutPrefix(child.ip, ip)
		}
		if longestPrefix(child.ip, ip) {
			return child, false, true, ip
		}

	}
	return &Node{}, false, false, ""
}

func (n *Node) Initiate(ip string) {
	n.ip = ip
	n.isLeaf = true
	n.children = []*Node{}
}

func main() {
	n := Node{}
	n.Add("192.168.0.1")
	n.Add("192.168.0.1")
	n.Add("192.168.0.2")
	n.Add("10.10.30.40")
	n.Add("250.10.30.40")
	fmt.Println("Does exists", n.exists("10.10.30.40"))
}
