package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"eval"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Println("Expression:")
	stdin.Scan()
	exprStr := stdin.Text()
	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Println("bad expression", err)
		os.Exit(1)
	}

	fmt.Println("Variables (e.g. x=3 y=4):")
	stdin.Scan()
	envStr := stdin.Text()

	env := eval.Env{}
	pairs := strings.Fields(envStr)
	for _, pair := range pairs {
		fields := strings.Split(pair, "=")
		if len(fields) != 2 {
			fmt.Println("invalid input:", pair)
			os.Exit(1)
		}

		varStr, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Println("bad value for", varStr)
			os.Exit(1)
		}

		env[eval.Var(varStr)] = val
	}

	fmt.Println(expr.Eval(env))
}
