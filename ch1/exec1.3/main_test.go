package main

import "testing"

var ss = []string{
	"apple",
	"boy",
	"cat",
	"dog",
	"eagle",
	"fox",
	"goose",
	"hawk",
	"ink",
	"jelly",
}

func BenchmarkSlowJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowJoin(ss)
	}
}

func BenchmarkFastJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastJoin(ss)
	}
}
