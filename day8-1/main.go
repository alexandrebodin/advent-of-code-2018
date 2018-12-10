package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	children []node
	metadata []int
}

func main() {
	fi, err := os.Open("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer fi.Close()

	numbers := []int{}

	s := bufio.NewScanner(fi)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		var i int
		fmt.Sscanf(s.Text(), "%d", &i)
		numbers = append(numbers, i)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	root, _ := parseNode(numbers)

	fmt.Println(sumMetadata(root))
}

func sumMetadata(n node) int {
	sum := 0
	for _, m := range n.metadata {
		sum += m
	}

	for _, child := range n.children {
		sum += sumMetadata(child)
	}

	return sum
}

func parseNode(arr []int) (node, int) {
	childrenCount, metaCount := arr[0], arr[1]
	children := []node{}
	l := 2

	for i := 0; i < childrenCount; i++ {
		node, len := parseNode(arr[l:])
		children = append(children, node)
		l += len
	}

	meta := arr[l : l+metaCount]

	return node{
		children: children,
		metadata: meta,
	}, l + metaCount
}
