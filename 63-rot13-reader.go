package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)
    for i := 0; i < n; i++ {
        p[i] = r.rot13(p[i])
    }
    return
}

func (r rot13Reader) rot13(x byte) byte {
    switch {
    case 'A' <= x && x <= 'Z':
        x = (x - 'A' + 13) % 26 + 'A'
    case 'a' <= x && x <= 'z':
        x = (x - 'a' + 13) % 26 + 'a'
    }
    return x
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
