package strings

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	s := []byte("가나다")
	s[2]++
	fmt.Println(string(s))
	// Output:
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

// To make sure not to know the result in compile time.
var s4 = time.Now().Format("20060102")

func BenchmarkSprintf4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s", s1, s2, s3, s4)
	}
}

func BenchmarkPlus4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = s1 + s2 + s3 + s4
	}
}

func BenchmarkSprint4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(s1, s2, s3, s4)
	}
}

func BenchmarkJoin4(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s1, s2, s3, s4}, "")
	}
}

func BenchmarkBytes(b *testing.B) {
	s1 := "hello"
	s2 := " world"
	s3 := " and and"
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString(s1)
		buf.WriteString(s2)
		buf.WriteString(s3)
		buf.WriteString(s4)
		_ = buf.String()
	}
}
