package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func Sha256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	s := strings.TrimSpace(os.Args[1])
	fmt.Print(Sha256(s))
}
