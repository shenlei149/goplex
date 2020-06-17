package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", float64(l))
}

func (u unary) String() string {
	return fmt.Sprintf("(%s %s)", string(u.op), u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("pow(%s, %s)", c.args[0].String(), c.args[1].String())
	case "sin":
		return fmt.Sprintf("sin(%s)", c.args[0].String())
	case "sqrt":
		return fmt.Sprintf("sqrt(%s)", c.args[0].String())
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
