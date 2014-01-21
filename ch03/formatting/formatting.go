package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", intForBool(true), intForBool(false))
	t := parseBool("True")
	f := parseBool("FALSE")
	fmt.Printf("%t %t\n", t, f)
	fmt.Println()
	
	fmt.Println("|123456" + strings.Repeat("|123456789", 5) + strings.Repeat("|1234567", 2) + "|")
	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|% 9b|% 7b|% 7b|\n", 37, 37, 37, 37, 37, -37, 37, -37)
}

func intForBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func parseBool(s string) bool {
	val, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Printf("cannot parse \"%s\"\n", s)
	}
	return val
}
