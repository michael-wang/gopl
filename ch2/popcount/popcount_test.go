package popcount

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestPopCount(t *testing.T) {
	check := func(x uint64, exp int, count func(uint64) int) {
		fmt.Printf("%s(%X)\n", varName(count), x)

		got := PopCount(x)
		if got != exp {
			t.Errorf("x: %64b, expected: %d, got: %d", x, exp, got)
		}
	}

	targets := []func(uint64) int{
		PopCount,
		PopCountByLoop,
		PopCountByShifting,
		PopCountByRightMostOne,
	}
	for _, count := range targets {
		check(1, 1, count)
		check(1<<63, 1, count)
		check(2, 1, count)
		check(3, 2, count)
		check(4, 1, count)
	}
}

func varName(f interface{}) string {
	fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return fname[strings.LastIndex(fname, ".")+1:]
}

func BenchmarkByTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(uint64(i))
	}
}

func BenchmarkByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(uint64(i))
	}
}

func BenchmarkByRightMostOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByRightMostOne(uint64(i))
	}
}
