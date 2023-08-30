package lesson_twenty_five

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()         // lock the account so we can get the balance
	defer a.lock.Unlock() // unlock the account after the balance has been retrieved
	return a.balance
}

func (a *Account) Deposit(amount int) {
	a.lock.Lock()         // lock the account so no one else can access the balance
	defer a.lock.Unlock() // unlock the account after the balance has been updated
	a.balance += amount
}

func (a *Account) Withdraw(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if amount > a.balance {
		fmt.Println("Not enough funds to withdraw")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n",
			amount, a.balance)
		a.balance -= amount
	}
}

func Start() string {
	fmt.Println("Lesson Twenty Five Started...")

	var acct Account
	acct.balance = 100
	fmt.Println("Balance :", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

	return "Lesson Twenty Five Complete!"
}
