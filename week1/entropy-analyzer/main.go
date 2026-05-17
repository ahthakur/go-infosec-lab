package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage - Entropy Analyzer <file>")
		os.Exit(1)
	}
	filename := os.Args[1]
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Invalid file ", filename)
		os.Exit(1)
	}
	filebytes := make(map[byte]int)

	for i := range contents {
		filebytes[contents[i]]++
	}
	totalnumberofbytes := len(contents)
	entropy := 0.0
	for _, count := range filebytes {
		p := float64(count) / float64(totalnumberofbytes)
		entropy += p * math.Log2(p)
	}
	entropy = -entropy
	fmt.Printf("File: %s\n", filename)
	fmt.Printf("Entropy: %.2f\n", entropy)

	if entropy < 5.0 {
		fmt.Println("Verdict: LOW (plaintext likely)")
	} else if entropy < 7.0 {
		fmt.Println("Verdict: MEDIUM (structured data)")
	} else {
		fmt.Println("Verdict: HIGH (encrypted/compressed/packed)")
	}
}
