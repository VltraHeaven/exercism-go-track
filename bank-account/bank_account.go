package account

import "sync"

type Account struct {
	balance int64
	active  bool
	mutex   sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, active: true}
}

func (a *Account) Balance() (balance int64, status bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.active {
		return
	}
	balance, status = a.balance, a.active
	return
}

func (a *Account) Deposit(amount int64) (balance int64, status bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	switch diff := a.balance + amount; {
	case diff < 0 || !a.active:
	default:
		a.balance += amount
		balance, status = a.balance, a.active
	}
	return
}

func (a *Account) Close() (payout int64, status bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.active {
		return
	}
	payout, status = a.balance, true
	a.balance, a.active = 0, false
	return
}
