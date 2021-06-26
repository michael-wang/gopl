package exec2_4

import (
	"testing"
)

func BenchmarkExec2_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
