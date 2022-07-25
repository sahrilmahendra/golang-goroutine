package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Deadlock Simulation
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1 ", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 ", user2.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Sahril",
		Balance: 200000,
	}

	user2 := UserBalance{
		Name:    "Mahendra",
		Balance: 500000,
	}

	go Transfer(&user1, &user2, 20000)
	go Transfer(&user2, &user1, 100000)

	time.Sleep(10 * time.Second)

	fmt.Println("user", user1.Name, "balance", user1.Balance)
	fmt.Println("user", user2.Name, "balance", user2.Balance)

}
