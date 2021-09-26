package main

import (
	"io"
	"os"
)

// file to open should be provided as an argument using os.Args, which is a []string

func main() {
	fp := os.Args[1]
	f, err := os.Open(fp)
	if err != nil {
		os.Exit(1)
	}
	//b1 := make([]byte, 99999)
	//n1, _ := f.Read(b1)
	//fmt.Printf("%d bytes:\n%s", n1, string(b1[:n1]))
	// Alternatively:
	io.Copy(os.Stdout, f)
}
