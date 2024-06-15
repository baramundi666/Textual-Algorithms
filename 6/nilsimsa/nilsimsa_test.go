package nilsimsa

import (
	"fmt"
	"os"
	"testing"
)

func TestNilsimsa(t *testing.T) {
	s1 := "To niedźwiedź czy może dźwiedź? Chyba nie dźwiedź."
	s2 := "Czy to dźwiedź, czy niedźwiedź? Może nie dźwiedź."
	s3 := "Najgłupsze zwierzę w dżungli? Niedźwiedź polarny."
	data := []struct {
		s1, s2 string
		want   int
	}{
		{s1, s2, 47},
		{s1, s3, 82},
		{s2, s3, 83},
	}
	for _, d := range data {
		if got := HammingDistance(Nilsimsa(d.s1), Nilsimsa(d.s2)); got != d.want {
			t.Errorf("HammingDistance(Nilsimsa(%v), Nilsimsa(%v)) == %d, want %d",
				d.s1, d.s2, got, d.want)
		}
	}
}

func TestNilsimsa2(t *testing.T) {
	s1, err := os.ReadFile("W1.txt")
	s2, err := os.ReadFile("W2.txt")
	s3, err := os.ReadFile("W3.txt")
	s4, err := os.ReadFile("W4.txt")
	s5, err := os.ReadFile("W5.txt")
	s6, err := os.ReadFile("W6.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	got := []int{}

	got = append(got, HammingDistance(Nilsimsa(string(s1)), Nilsimsa(string(s2))))
	got = append(got, HammingDistance(Nilsimsa(string(s1)), Nilsimsa(string(s3))))
	got = append(got, HammingDistance(Nilsimsa(string(s1)), Nilsimsa(string(s4))))
	got = append(got, HammingDistance(Nilsimsa(string(s1)), Nilsimsa(string(s5))))
	got = append(got, HammingDistance(Nilsimsa(string(s1)), Nilsimsa(string(s6))))

	fmt.Printf("Results: %v\n", got)
}
