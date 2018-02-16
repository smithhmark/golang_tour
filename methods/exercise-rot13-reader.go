package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rdr rot13Reader) Read(b []byte) (int, error) {
        cnt, err := rdr.r.Read(b)
        if err != nil {
                return cnt, err
        }

        bottom_lower := int('a')
        top_lower := int('z')
        bottom_upper := int('A')
        top_upper := int('Z')
        for ii := 0 ; ii < len(b) ; ii++ {
                ord := int(b[ii])
                if bottom_lower <= ord  && ord <= top_lower{
                        tmp := (ord - bottom_lower + 13) % 26 + bottom_lower
                        b[ii] = byte(tmp)
                } else if bottom_upper <= ord  && ord <= top_upper {
                        tmp := (ord - bottom_upper + 13) % 26 + bottom_upper
                        b[ii] = byte(tmp)
                }

        }
        return cnt, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
