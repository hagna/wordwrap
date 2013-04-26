package wordwrap

import (
	"bytes"
	_ "fmt"
	"io"
)

func findMark(data []byte, wrap int) (res int) {
	res = wrap
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

// Delete up to n old bytes from s
// if n < 0 delete them all
func rmByte(s, old []byte, n int) []byte {
	m := bytes.Index(s, old)
	res := s[:]
	c := 0
	for m != -1 {
		res = append(res[:m], res[m+1:]...)
		c++
		if c == n {
			break
		}
	}
	return res
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
*/
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
			s = bytes.Replace(s, []byte{'\n'}, nil, -1)
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
