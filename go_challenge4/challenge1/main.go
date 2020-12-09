package main

import (
	"fmt"

	"github.com/mchaker/golang_coding_challenges/go_challenge4/challenge1/bank"
)

func main() {
	var myAccount = new(bank.BankAccount)

	fmt.Println("=== Challenge 1 ===")

	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("--> Depositing $100.00")
	myAccount.Deposit("100")
	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("<-- Withdrawing $45.67")
	myAccount.Withdraw("45.67")
	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("--> Depositing $0.99")
	myAccount.Deposit("0.99")
	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("--> Depositing $0.00")
	myAccount.Deposit("0.00")
	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("--- This is an imaginary kind bank that only acts on amounts > $0.005 :)")

	fmt.Println("--> Depositing $0.001")
	myAccount.Deposit("0.001")
	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("--> Depositing $0.009")
	myAccount.Deposit("0.009")
	fmt.Println("Account balance: ", myAccount.CheckBalance())

	fmt.Println("<-- Withdrawing $0.001")
	myAccount.Withdraw("0.001")
	fmt.Println("Account balance: ", myAccount.CheckBalance())
}
