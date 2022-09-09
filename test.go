package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "21. fasdfas"
	index := strings.Index(s, ".")
	fmt.Println(s[:index])
}
