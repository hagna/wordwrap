package wordwrap

import (
	"bytes"
	"io"
	"fmt"
	"bufio"
)

/*
	If all the text has newlines around 80 then return 80 because
	it's wrapped to that.
*/
func findSoftwrap(r io.Reader) (res int) {
	s := bufio.NewScanner(r)
	for {
		ok := s.Scan()
		if !ok {
			if s.Err == nil {
				break
			}
		}
	}
	return 80 // TODO 	

}

func findMark(data []byte, wrap int) (res int) {
	res = -1
	if len(data) > wrap {
		if data[wrap] == ' ' {
			res = wrap
		} else {
			if len(data) > wrap+1 {
				if data[wrap+1] == ' ' {
					res = wrap
					return
				}
			}
			res = bytes.LastIndex(data[:wrap+1], []byte{' '})
			if res == -1 {
				res = wrap
			}
		}
	}
	return
}

// Wrapit(r, 4, cb) would call cb at or before every
// 4th byte consecutive byte not including space.
// For example
/* Wrapit(strings.NewReader("this would wrap. like this", 6, cb) would call cb on each of these:
this
would
wrap.
like
this
func Wrapit(r io.Reader, wrap int, cb func([]byte)) {
	frag := []byte{}
loop:
	for {
		s := []byte{}
		for len(s) < wrap+2 {
			buf := make([]byte, (wrap+2)-len(frag)-len(s))
			n, err := io.ReadFull(r, buf)
			buf = append(frag, buf[:n]...)
			frag = []byte{}
			s = append(s, buf...)
			//s = bytes.Replace(s, []byte{'\n'}, nil, -1)
			if err != nil {
				cb(s)
				break loop
			}

		}
		mark := findMark(s, wrap)
		frag = append(frag, s[mark+1:]...)
		cbthis := []byte{}
		cbthis = append(cbthis, s[:mark+1]...)
		s = []byte{}
		cb(cbthis)
		// a[:i+1] and a[i+1:] is confusing because the first includes i, but not i+1
		// while the second includes i+1
	}

}
*/
/*
It's an easier problem when softreturns aren't confused with hard returns, and 
we do that by making the wordprocessor, madman in this case, output one long line
for sentences that are meant for the same paragraph.
*/
func Wrapit(r io.Reader, wrap int, cb func([]byte)) {
	s := bufio.NewScanner(r) 
	for {
		ok := s.Scan()
		if !ok {
			break
		}
		all := s.Bytes()
		z := all
		for {
			// chomp blanks
			j := 0
			for z[j] == ' ' {
				j++
			}
			z = z[j:]
			if len(z) <= wrap {
				cb(z)
				break 
			}
			mark := findMark(z, wrap)
			if mark == -1 {
				fmt.Println("shouldn't get here: mark is -1 which means len(data) < wrap")
			}
			cbthis := []byte{}
			
			cbthis = append(cbthis, z[:mark+1]...)
			cb(cbthis)
			z = z[mark+1:]
		}
	}
}
		
