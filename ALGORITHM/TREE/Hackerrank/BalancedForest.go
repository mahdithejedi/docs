package main

import (
	"fmt"
)

/*
 * Complete the 'balancedForest' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY c
 *  2. 2D_INTEGER_ARRAY edges
 */

type Node struct {
	value  int32
	number int32
	nodes  []*Node
}

func (n *Node) SearchByNumber(number int32) *Node {
	if n.number == number {
		return n
	}
	for _, child := range n.nodes {
		child.SearchByNumber(number)
	}
	return nil
}

func (n *Node) InsertByNumber(parent_node_number int32, number int32, value int32) bool {
	if n == nil {
		n.value = value
		n.number = number
	}
	parent_node := n.SearchByNumber(parent_node_number)
	if parent_node == nil {
		return false
	}
	parent_node.nodes = append(parent_node.nodes, &Node{value: value, number: number})
	return true
}

func balancedForest(c []int32, edges [][]int32) *Node {
	node := &Node{
		value:  c[0],
		number: 1,
	}
	values := c[1:]
	for index, _ := range values {
		node.InsertByNumber(
			edges[index][0], edges[index][1], c[index],
		)
	}
	return node
}

func main() {
	edges := [][]int32{
		{1, 2},
		{1, 3},
		{3, 5},
		{1, 4},
	}
	values := []int32{
		1, 2, 2, 1, 1,
	}
	res := balancedForest(values, edges)
	fmt.Print(res.nodes[1].nodes)
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	//
	//qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//q := int32(qTemp)
	//
	//for qItr := 0; qItr < int(q); qItr++ {
	//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//	checkError(err)
	//	n := int32(nTemp)
	//
	//	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	//
	//	var c []int32
	//
	//	for i := 0; i < int(n); i++ {
	//		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
	//		checkError(err)
	//		cItem := int32(cItemTemp)
	//		c = append(c, cItem)
	//	}
	//
	//	var edges [][]int32
	//	for i := 0; i < int(n)-1; i++ {
	//		edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")
	//
	//		var edgesRow []int32
	//		for _, edgesRowItem := range edgesRowTemp {
	//			ed(edgesRow) != 2 {
	//			panic("Bad input")
	//		}
	//
	//		edges = append(edges, edgesRow)
	//	}
	//
	//	result := balancedForest(c, edges)
	//
	//	fmt.Fprintf(writer, "%d\n", result)
	//}
	//
	////writer.Flush()gesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
	//			checkError(err)
	//			edgesItem := int32(edgesItemTemp)
	//			edgesRow = append(edgesRow, edgesItem)
	//		}
	//
	//		if len(edgesRow) != 2 {
	//			panic("Bad input")
	//		}
	//
	//		edges = append(edges, edgesRow)
	//	}
	//
	//	result := balancedForest(c, edges)
	//
	//	fmt.Fprintf(writer, "%d\n", result)
	//}
	//
	//writer.Flush()
}

//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
