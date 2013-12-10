package main

import (
	"github.com/hagna/wordwrap"
	"bufio"
	"os"
	"fmt"
	"flag"
)

var wrapat = flag.Int("w", 70, "default column to wrap")


func IsFile(fname string) bool {
	s, err := os.Stat(fname)
	if err == nil {
		if s.IsDir() == false {
			return true
		}
	}
	return false
}




func main() {
	flag.Parse()
	usestdin := true
	var bi *bufio.Reader
	bw := bufio.NewWriter(os.Stdout)
	for _, fname := range flag.Args() {
		if IsFile(fname) {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Println(err)
				continue
			}
			usestdin = false
			bi := bufio.NewReader(f)
			wordwrap.Wrapit(bi, *wrapat, func(s []byte) {
				fmt.Fprintln(bw, string(s))
			})
		}
	}
	if usestdin {
		bi = bufio.NewReader(os.Stdin)
		wordwrap.Wrapit(bi, *wrapat, func(s []byte) {
			fmt.Fprintln(bw, string(s))
		})
	}
	bw.Flush()
}
