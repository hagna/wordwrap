package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
)

func IsFile(fname string) bool {
	s, err := os.Stat(fname)
	if err == nil {
		if s.IsDir() == false {
			return true
		}
	}
	return false
}



func padit(n int, o io.Writer, i io.Reader) {
	s := bufio.NewScanner(i)	
	for {
		ok := s.Scan()
		if !ok {
			if s.Err() == nil {
				break
			}
		}
		for j := 0; j < n; j++ {
			fmt.Fprint(o, " ")
		}
		fmt.Fprintln(o, s.Text())
	}
}
			
		

func main() {
	usestdin := true
	var bi *bufio.Reader
	bw := bufio.NewWriter(os.Stdout)
	for _, fname := range os.Args[1:] {
		if IsFile(fname) {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Println(err)
				continue
			}
			usestdin = false
			bi := bufio.NewReader(f)
			padit(10, bw, bi)
		}
	}
	if usestdin {
		bi = bufio.NewReader(os.Stdin)
		padit(10, bw, bi)
	}
	bw.Flush()
}
