package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

func main() {

	algo := flag.String("algo", "md5", "Hash algorithm: md5 or sha256")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Usage : Provide Algo flag hash and filename ")
		os.Exit(1)
	}

	userHash := args[0]
	fileName := args[1]
	fmt.Println("Algorithm:", *algo)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found : ", err)
		os.Exit(1)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fileHash := "sample"
	for scanner.Scan() {
		word := scanner.Text()
		if *algo == `md5` {
			fileHash = fmt.Sprintf("%x", md5.Sum([]byte(word)))
		} else if *algo == `sha256` {
			fileHash = fmt.Sprintf("%x", sha256.Sum256([]byte(word)))
		}
		if fileHash == userHash {
			fmt.Println("Cracked :-> ", word)
		}
	}
}
