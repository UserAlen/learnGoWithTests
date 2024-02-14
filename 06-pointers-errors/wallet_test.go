package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance().String()
	/* fmt.Printf("address of balance in test is %p \n", &wallet.balance) */
	want := Bitcoin(10).String()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
