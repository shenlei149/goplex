package bank

type withdrawArg struct {
	amount int
	ret    chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan withdrawArg)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	ret := make(chan bool)
	withdraw <- withdrawArg{amount, ret}
	return <-ret
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdrawArg := <-withdraw:
			if balance >= withdrawArg.amount {
				balance -= withdrawArg.amount
				withdrawArg.ret <- true
			} else {
				withdrawArg.ret <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
