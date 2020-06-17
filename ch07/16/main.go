package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"eval"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exprStr := r.FormValue("expr")
		if exprStr == "" {
			http.Error(w, "no expression", http.StatusBadRequest)
			return
		}

		expr, err := eval.Parse(exprStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("bad expression: %s. Err: %s", exprStr, err), http.StatusBadRequest)
			return
		}

		envStr := r.FormValue("env")
		env := eval.Env{}
		pairs := strings.Fields(envStr)
		for _, pair := range pairs {
			fields := strings.Split(pair, "=")
			if len(fields) != 2 {
				http.Error(w, fmt.Sprintf("invalid input: %s", pair), http.StatusBadRequest)
				return
			}

			varStr, valStr := fields[0], fields[1]
			val, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("bad value for %s", valStr), http.StatusBadRequest)
				return
			}

			env[eval.Var(varStr)] = val
		}

		fmt.Fprintln(w, expr.Eval(env))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
