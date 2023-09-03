package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY stringList
 *  2. STRING_ARRAY queries
 */

// __URL__ = https://www.hackerrank.com/challenges/sparse-arrays/problem
// __DS__

func getHash(str string) string {
	//h := fnv.New32a()
	//h.Write([]byte(str))
	//return h.Sum32()
	return str
}

type hashValue struct {
	count uint32
	idx   int
}

func getHashes(queries *[]string) *map[string]hashValue {
	hashes := make(map[string]hashValue)
	for idx, value := range *queries {
		hashes[getHash(value)] = hashValue{
			0,
			idx,
		}
	}
	return &hashes
}

func matchingStrings(stringList []string, queries []string) []int32 {
	hashes := getHashes(&queries)
	for _, value := range stringList {
		_localHash, exists := (*hashes)[getHash(value)]
		if exists == false {
			continue
		}
		(*hashes)[getHash(value)] = hashValue{
			_localHash.count + 1,
			_localHash.idx,
		}
	}
	values := make([]int32, len(queries))
	for _, value := range *hashes {
		fmt.Println(values, value.idx)
		values[value.idx] = int32(value.count)
	}
	return values
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stringListCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var stringList []string

	for i := 0; i < int(stringListCount); i++ {
		stringListItem := readLine(reader)
		stringList = append(stringList, stringListItem)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(stringList, queries)

	for i, resItem := range res {
		fmt.Fprintf(os.Stdout, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(os.Stdout, "\n")
		}
	}

	//fmt.Fprintf(writer, "\n")
	//
	//writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Best golang solutions
//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//)
//
//var N, Q int
//
//var strings, queries []string
//
//func main() {
//
//	io := bufio.NewReader(os.Stdin)
//
//	fmt.Fscan(io, &N)
//	strings = make([]string, N)
//	for i := 0; i < N; i++ {
//		fmt.Fscan(io, &strings[i])
//	}
//
//	fmt.Fscan(io, &Q)
//	queries = make([]string, Q)
//	for i := 0; i < Q; i++ {
//		fmt.Fscan(io, &queries[i])
//	}
//
//	for i := 0; i < len(queries); i++ {
//		var s int
//		for j := 0; j < len(strings); j++ {
//			if queries[i] == strings[j] {
//				s += 1
//			}
//		}
//		fmt.Println(s)
//	}
//}
