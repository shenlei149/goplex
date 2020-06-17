package word

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		nn := rng.Intn(0x1000) // random rune up to '\u0999'
		var r rune
		if nn%10 == 0 {
			r = ' '
		} else if nn%5 == 0 {
			r = []rune{',', ';', '=', '-', '?'}[nn%5]
		} else {
			r = rune(nn)
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 10 // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(26) + 'A') // random rune up to \u99
		runes[i] = r
		runes[n-1-i] = rune(r + 1)
	}
	return "?" + string(runes) + "!"
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = true", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
