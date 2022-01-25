package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Give me some integers with space in between->")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	elements := strings.Split(strings.TrimSpace(line), " ")

	numbers:=[]int{}
	for _, value := range elements {
		result,_ := strconv.Atoi(value)
		// numbers[index]=result
		numbers=append(numbers,result)
	}

	sort.Ints(numbers)
	fmt.Println("\nSorted numbers are:", numbers)

}
