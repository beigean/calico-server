package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r *MyReader) Read(b []byte) (n int, err error) {
	for n, err = 0, nil; n < len(b); n++ {
		b[n] = 'A'
	}
	return n, err
}

func main() {
	reader.Validate(&MyReader{})

	// r := MyReader{}
	// b := make([]byte, 8)
	// for {
	// 	n, err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
}
