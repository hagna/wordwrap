package wordwrap

import "testing"
import "strings"

/*func TestWrap(t *testing.T) {
	b := make([]byte, 1000)
	for i:=0 ; i<1000; i++ {
		b[i] = 'c'
	}
	s := strings.NewReader(string(b))
	wr, err := NewReader(s, 10)
	if err != nil {
		t.Fatalf(err.Error())
	}

	s2, err := ioutil.ReadAll(wr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	f := strings.Split(string(s2), "\n")
	if len(f[0]) != 10 {
		t.Errorf("should have been 10 long instead got %s", f[0])
	}

}
*/

func TestWrapSimple(t *testing.T) {
	s := "this would wrap. like this"
	a := strings.NewReader(s)
	b := []byte{}
	right := []byte(`this 
would 
wrap. 
like 
this
`)
	cb := func(t []byte) {
		b = append(b, append(t, '\n')...)
	}
	Wrapit(a, 6, cb)
	if len(b) != len(right) {
		t.Errorf("lengths don't match it is %d but it should be %d\n", len(b), len(right))
		t.Errorf("can you spot the difference\n++>%s<--\n-->%s<--\n", string(right), string(b))
	}
	for i := 0; i < len(right); i++ {
		if right[i] != b[i] {
			t.Errorf("mismatch at character %d should be \"%s\" but it is \"%s\"\n", i, right[:i], b[:i])
		}
	}
}

func TestFindMark(t *testing.T) {
	a := []byte("0123456789 ")
	m := findMark(a, 10)
	if m != 10 {
		t.Errorf("%s should have wrapped at 10 and not %d\n", string(a), m)
	}
	a = []byte("0123456789A ")
	m = findMark(a, 10)
	if m != 10 {
		t.Errorf("%s should have wrapped at 10 and not %d\n", string(a), m)
	}
	a = []byte("01234567 9A")
	m = findMark(a, 10)
	if m != 8 {
		t.Errorf("%s should have wrapped at 8 and not %d\n", string(a), m)
	}
	a = []byte("0123456789A")
	m = findMark(a, 10)
	if m != 10 {
		t.Errorf("%s should have wrapped at 10 and not %d\n", string(a), m)
	}

}
