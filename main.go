package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("/home/jtfogar/go/src/github.com/jtfogarty/DocuRe/bible-data/bibles.txt")
	check(err)

	f, err 

}
