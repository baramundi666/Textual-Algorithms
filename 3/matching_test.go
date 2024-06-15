package matching

import (
	"os"
	"slices"
	"testing"
)

func TestPreprocess(t *testing.T) {
	data := []string{
		"aaaaaaa",
		"pies",
		"dźwiedź",
		"owocowo",
		"indianin",
		"nienapełnienie",
	}
	for _, in := range data {
		got := Preprocess([]byte(in))
		want := SimplePreprocess([]byte(in))
		if !slices.Equal(got, want) {
			t.Errorf(`Preprocess(%#v) == %#v want %#v`,
				in, got, want)
		}
	}
}

func indices(pat, text []byte) []int {
	r := []int{}
	for i := 0; i+len(pat) <= len(text); i++ {
		if slices.Equal(text[i:i+len(pat)], pat) {
			r = append(r, i)
		}
	}
	return r
}

func TestNaive(t *testing.T) {
	pat := []byte("chmur")
	text := []byte("Pochmurne chmury")

	got := []int{}
	Naive(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf("error!")
	}
}

func TestBackwardNaive(t *testing.T) {
	pat := []byte("chmur")
	text := []byte("Pochmurne chmury")

	got := []int{}
	BackwardNaive(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf("error!")
	}
}

func TestBoyerMoore(t *testing.T) {
	pat := []byte("chmur")
	text := []byte("Pochmurne chmury")

	got := []int{}
	BoyerMoore(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf("error!")
	}
}

func TestKMP(t *testing.T) {
	pat := []byte("chmur")
	text := []byte("Pochmurne chmury")

	got := []int{}
	KMP(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf("error!")
	}
}

func TestKarpRabin(t *testing.T) {
	pat := []byte("chmur")
	text := []byte("Pochmurne chmury")

	got := []int{}
	KarpRabin(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf("error!")
	}
}

func TestShiftOr(t *testing.T) {
	pat := []byte("chmur")
	text := []byte("Pochmurne chmury")

	got := []int{}
	ShiftOr(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf("error!")
	}
}

func BenchmarkShortNaive(b *testing.B) {
	pat := []byte("chmura")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkShortBackwardNaive(b *testing.B) {
	pat := []byte("chmur")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkShortBoyerMoore(b *testing.B) {
	pat := []byte("chmur")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkShortKMP(b *testing.B) {
	pat := []byte("chmur")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

// func BenchmarkShortKarpRabin(b *testing.B) {
// 	pat := []byte("chmur")
// 	text, _ := os.ReadFile("kordian.txt")

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		KarpRabin(pat, text, func(int) {})
// 	}
// }

func BenchmarkShortShiftOr(b *testing.B) {
	pat := []byte("chmur")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}

func BenchmarkLongNaive(b *testing.B) {
	pat := []byte("katedralnego")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkLongBackwardNaive(b *testing.B) {
	pat := []byte("katedralnego")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkLongBoyerMoore(b *testing.B) {
	pat := []byte("katedralnego")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkLongKMP(b *testing.B) {
	pat := []byte("katedralnego")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

// func BenchmarkLongKarpRabin(b *testing.B) {
// 	pat := []byte("katedralnego")
// 	text, _ := os.ReadFile("kordian.txt")

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		KarpRabin(pat, text, func(int) {})
// 	}
// }

func BenchmarkLongShiftOr(b *testing.B) {
	pat := []byte("katedralnego")
	text, _ := os.ReadFile("kordian.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}
