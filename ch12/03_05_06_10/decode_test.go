package sexpr

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type account struct {
		Changed bool
		num     float64
		name    string
	}

	b, err := Marshal(account{true, 300.9, "le"})
	if err != nil {
		t.Fatalf("Marshal failed: %v\n", err)
	}
	var a account
	err = Unmarshal(b, &a)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v\n", err)
	}
	fmt.Println(a)
}
