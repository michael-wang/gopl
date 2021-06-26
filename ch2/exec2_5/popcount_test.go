package exec2_5

import (
	"testing"
)

func BenchmarkExec2_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
