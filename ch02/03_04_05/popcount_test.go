package popcount

import "testing"

// go version go1.14.2 linux/amd64
// goos: linux
// goarch: amd64
// BenchmarkPopCount-12                    1000000000               0.306 ns/op
// BenchmarkPopCountByLoop-12              69560022                16.1 ns/op
// BenchmarkPopCountByClearing-12          53670075                20.7 ns/op
// BenchmarkPopCountByShifting-12          21200104                54.5 ns/op

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// Table vs. Clearing
// BenchmarkPopCount1-12                           1000000000               0.558 ns/op
// BenchmarkPopCount10-12                          364981153                3.16 ns/op
// BenchmarkPopCount100-12                         35675245                32.7 ns/op
// BenchmarkPopCount1000-12                         4383252               276 ns/op
// BenchmarkPopCount10000-12                         438396              2811 ns/op
// BenchmarkPopCount100000-12                         37890             27782 ns/op
// BenchmarkPopCountByClearing1-12                 43746872                26.2 ns/op
// BenchmarkPopCountByClearing10-12                 4733960               216 ns/op
// BenchmarkPopCountByClearing100-12                 595908              2223 ns/op
// BenchmarkPopCountByClearing1000-12                 51096             21526 ns/op
// BenchmarkPopCountByClearing10000-12                 5390            203897 ns/op
// BenchmarkPopCountByClearing100000-12                 525           2134227 ns/op

func benchmarkPopCount(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			PopCount(0x1234567890ABCDEF)
		}
	}
}

func BenchmarkPopCount1(b *testing.B)      { benchmarkPopCount(b, 1) }
func BenchmarkPopCount10(b *testing.B)     { benchmarkPopCount(b, 10) }
func BenchmarkPopCount100(b *testing.B)    { benchmarkPopCount(b, 100) }
func BenchmarkPopCount1000(b *testing.B)   { benchmarkPopCount(b, 1000) }
func BenchmarkPopCount10000(b *testing.B)  { benchmarkPopCount(b, 10000) }
func BenchmarkPopCount100000(b *testing.B) { benchmarkPopCount(b, 100000) }

func benchmarkPopCountByClearing(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			PopCountByClearing(0x1234567890ABCDEF)
		}
	}
}

func BenchmarkPopCountByClearing1(b *testing.B)      { benchmarkPopCountByClearing(b, 1) }
func BenchmarkPopCountByClearing10(b *testing.B)     { benchmarkPopCountByClearing(b, 10) }
func BenchmarkPopCountByClearing100(b *testing.B)    { benchmarkPopCountByClearing(b, 100) }
func BenchmarkPopCountByClearing1000(b *testing.B)   { benchmarkPopCountByClearing(b, 1000) }
func BenchmarkPopCountByClearing10000(b *testing.B)  { benchmarkPopCountByClearing(b, 10000) }
func BenchmarkPopCountByClearing100000(b *testing.B) { benchmarkPopCountByClearing(b, 100000) }
