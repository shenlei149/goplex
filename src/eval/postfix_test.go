package eval

import (
	"fmt"
	"testing"
)

func TestPostfix(t *testing.T) {
	// don't extend parser, so have to construct a syntaxtree
	expr := postfix{"++", call{"pow", []Expr{Var("x"), literal(3)}}}
	env := Env{"x": 12}
	got := fmt.Sprintf("%.6g", expr.Eval(env))
	if got != "1729" {
		t.Errorf("got %q, want %q\n", got, "1729")
	}
}
