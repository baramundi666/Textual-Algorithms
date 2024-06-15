package suffixarray_test

import (
	"bufio"
	//"github.com/MarcinCiura/AT-lab/5/suffixarray"
	"github.com/cloudflare/ahocorasick"
	"index/suffixarray"
	"io/ioutil"
	"os"
	"testing"
)

type el struct {
	in   string
	want []int
}

type Index struct {
	suffixes []int
	text     []byte
}

func readFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func readLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ananas
//func TestSuffixArray(t *testing.T) {
//	ind := suffixarray.New([]byte("ananas"))
//	data := []el{
//		{
//			"nan",
//			[]int{1},
//		},
//		{
//			"na",
//			[]int{1, 3},
//		},
//		{
//			"a",
//			[]int{0, 2, 4},
//		},
//		{
//			"ka",
//			[]int{},
//		},
//	}
//	for _, d := range data {
//		//got := ind.Lookup([]byte(d.in), 6)
//		got := ind.LookupAll([]byte(d.in))
//		if fmt.Sprintf("%#v", got) != fmt.Sprintf("%#v", d.want) {
//			t.Errorf("Split(%#v) == %#v want %#v",
//				d.in, got, d.want)
//		}
//
//	}
//
//}

func BenchmarkAhoCorasickBuild(b *testing.B) {
	geny, err := readLines("geny.txt")
	if err != nil {
		b.Fatalf("Failed to read geny.txt: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher := ahocorasick.NewStringMatcher(geny)
		_ = matcher
	}
}

//func BenchmarkSimpleSuffixArrayBuild(b *testing.B) {
//	mDNA, err := readFile("mDNA.txt")
//	if err != nil {
//		b.Fatalf("Failed to read mDNA.txt: %v", err)
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		suffixArr := suffixarray.New([]byte(mDNA))
//		_ = suffixArr
//	}
//}
//
//
//func BenchmarkLibrarySuffixArrayBuild(b *testing.B) {
//	mDNA, err := readFile("mDNA.txt")
//	if err != nil {
//		b.Fatalf("Failed to read mDNA.txt: %v", err)
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		suffixArr := suffixarray.New([]byte(mDNA))
//		_ = suffixArr
//	}
//}

func BenchmarkAhoCorasickSearch(b *testing.B) {
	mDNA, err := readFile("mDNA.txt")
	if err != nil {
		b.Fatalf("Failed to read mDNA.txt: %v", err)
	}

	geny, err := readLines("geny.txt")
	if err != nil {
		b.Fatalf("Failed to read geny.txt: %v", err)
	}

	matcher := ahocorasick.NewStringMatcher(geny)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match([]byte(mDNA))
	}
}

//
//func BenchmarkSimpleSuffixArraySearch(b *testing.B) {
//	mDNA, err := readFile("mDNA.txt")
//	if err != nil {
//		b.Fatalf("Failed to read mDNA.txt: %v", err)
//	}
//
//	geny, err := readLines("geny.txt")
//	if err != nil {
//		b.Fatalf("Failed to read geny.txt: %v", err)
//	}
//
//	suffixArr := suffixarray.New([]byte(mDNA))
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for _, gene := range geny {
//			suffixArr.LookupAll([]byte(gene))
//		}
//	}
//}

//
func BenchmarkLibrarySuffixArraySearch(b *testing.B) {
	mDNA, err := readFile("mDNA.txt")
	if err != nil {
		b.Fatalf("Failed to read mDNA.txt: %v", err)
	}

	geny, err := readLines("geny.txt")
	if err != nil {
		b.Fatalf("Failed to read geny.txt: %v", err)
	}

	suffixArr := suffixarray.New([]byte(mDNA))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, gene := range geny {
			suffixArr.Lookup([]byte(gene), -1)
		}
	}
}
