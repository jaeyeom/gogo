package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("% x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
	// Output:
	// 0e ff 81 04 01 02 ff 82 00 01 0c 01 0c 00 00 08
	// ff 82 00 01 01 4e 01 4a
	// map[N:J]
}
