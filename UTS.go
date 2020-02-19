package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count++
		kata2 := strings.Split(scanner.Text()," ")
		m := make(map[string]int) 
		for _, kata := range kata2 { 
		   m[kata] += 1 
	   }
	   fmt.Println(kata2)
	   fmt.Println()
	   fmt.Println(m)
	}

}
