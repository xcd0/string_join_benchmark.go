package main

import (
	"bytes"
	"strings"
	"testing"
)

var m = [...]string{
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
	"AAAAAAAAA",
}

func BenchmarkStringsJoin____(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(m[:], ",") + ","
	}
}

func BenchmarkAppendOperator_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 string
		for _, v := range m {
			m2 += "," + v
		}
		m2 += ","
	}
}

func BenchmarkHardCoding_____(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m[0] + "," + m[1] + "," + m[2] + "," + m[3] + "," + m[4] + "," + m[5] + "," + m[6]
	}
}

func BenchmarkByteArray______(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 []byte
		for _, v := range m {
			m2 = append(m2, v...)
			m2 = append(m2, ',')
		}
		_ = string(m2)
	}
}

func BenchmarkCapByteArray___(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = make([]byte, 0, 100)
		for _, v := range m {
			m2 = append(m2, v...)
			m2 = append(m2, ',')
		}
		_ = string(m2)
	}
}

func BenchmarkBytesBuffer____(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 bytes.Buffer
		for _, v := range m {
			m2.Write([]byte(v))
			m2.Write([]byte{','})
		}
		_ = m2.String()
	}
}

func BenchmarkCapBytesBuffer_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = bytes.NewBuffer(make([]byte, 0, 100))
		for _, v := range m {
			m2.Write([]byte(v))
			m2.Write([]byte{','})
		}
		_ = m2.String()
	}
}

func BenchmarkCapBytesBuffer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = bytes.NewBuffer(make([]byte, 0, 100))
		for _, v := range m {
			m2.WriteString(v)
			m2.WriteString(",")
		}
		_ = m2.String()
	}
}
