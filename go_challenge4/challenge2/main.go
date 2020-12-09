package main

import (
	"fmt"

	"github.com/mchaker/golang_coding_challenges/go_challenge4/challenge2/bank"
)

func main() {
	userName := "John User"
	_, err := bank.MakeAccount(userName, "")

	if err != nil {
		fmt.Println("[ERROR] Unable to make bank account for ", userName, ".")
		fmt.Println("Error details: ", err)
	}

	fmt.Println("=== Challenge 2 ===")

	fmt.Println("----- Joint account demonstration")
	//joint accounts - make two separate accounts, show balances
	firstAccount, _ := bank.MakeAccount("Joe User", "2205.00")
	secondAccount, _ := bank.MakeAccount("Jane User", "9240.00")

	fmt.Println("Account 1 balance: $", firstAccount.CheckBalance())
	fmt.Println("Account 2 balance: $", secondAccount.CheckBalance())

	// now let's make a joint account
	// QUESTION: Why are the reference operators AND star operators
	// (in the JointAccount function definition/signature) needed for this to
	// work properly? Deposit and Withdraw worked with only the star operators
	// (which I also don't understand)
	jointAccount, _ := bank.JointAccount(&firstAccount, &secondAccount)
	fmt.Println("Joint account has been made.")

	fmt.Println("Joint account balance: $", jointAccount.CheckBalance())

	// and let's check the original accounts to make sure the money was successfully transferred
	fmt.Println("Account 1 balance: $", firstAccount.CheckBalance())
	fmt.Println("Account 2 balance: $", secondAccount.CheckBalance())

	// Joint account tests
	// Is there a better way to organize tests for a Go package, instead of
	// Println tests like these?
	fmt.Println("----- Account tests")
	//fmt.Println("Joint account starting balance: $", jointAccount.CheckBalance())

	var amount, testTitle string

	testAccount, _ := bank.MakeAccount("Test User", "0")

	testTitle = "=== Test: Deposit a positive amount with two decimal places"
	fmt.Println(testTitle)
	amount = "1.23"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit a positive amount with one decimal place"
	fmt.Println(testTitle)
	amount = "1.2"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit a positive amount with a decimal point and no decimal places"
	fmt.Println(testTitle)
	amount = "1."
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit a positive amount with only dollars and no decimal point"
	fmt.Println(testTitle)
	amount = "1"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit a positive amount with only cents and a zero dollar amount"
	fmt.Println(testTitle)
	amount = "0.23"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit a positive amount with only cents and no dollar amount"
	fmt.Println(testTitle)
	amount = ".23"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit a negative amount"
	fmt.Println(testTitle)
	amount = "-1.45"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Withdraw more than what is in the account"
	fmt.Println(testTitle)
	amount = "30.23"
	fmt.Println("Withdrawing ", amount)
	testAccount.Withdraw(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Withdraw a negative amount"
	fmt.Println(testTitle)
	amount = "-1.23"
	fmt.Println("Withdrawing ", amount)
	testAccount.Withdraw(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Deposit 0"
	fmt.Println(testTitle)
	amount = "0"
	fmt.Println("Depositing ", amount)
	testAccount.Deposit(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Withdraw 0"
	fmt.Println(testTitle)
	amount = "0"
	fmt.Println("Withdrawing ", amount)
	testAccount.Withdraw(amount)
	fmt.Println("→ Account balance: $", testAccount.CheckBalance())

	testTitle = "=== Test: Making an account with 0 balance"
	fmt.Println(testTitle)
	amount = "0"
	newAccount, _ := bank.MakeAccount("Joe User", amount)
	fmt.Println("→ Account balance: $", newAccount.CheckBalance())

	testTitle = "=== Test: Making an account with negative balance"
	fmt.Println(testTitle)
	amount = "-1.23"
	newAccount2, _ := bank.MakeAccount("Joe User2", amount)
	fmt.Println("→ Account balance: $", newAccount2.CheckBalance())

	testTitle = "=== Test: Make a joint account with 0 balance"
	fmt.Println(testTitle)
	amount = "0"
	newAccountA, _ := bank.MakeAccount("First User", amount)
	newAccountB, _ := bank.MakeAccount("Second User", amount)
	testJointAccount, _ := bank.JointAccount(&newAccountA, &newAccountB)
	fmt.Println("Account A Balance: $", newAccountA.CheckBalance())
	fmt.Println("Account B Balance: $", newAccountB.CheckBalance())
	fmt.Println("→ Joint Account Balance: $", testJointAccount.CheckBalance())

	testTitle = "=== Test: Make a joint account with a positive balance"
	fmt.Println(testTitle)
	amount = "1.23"
	newAccountA, _ = bank.MakeAccount("First User", amount)
	newAccountB, _ = bank.MakeAccount("Second User", amount)
	fmt.Println("Account A Balance: $", newAccountA.CheckBalance())
	fmt.Println("Account B Balance: $", newAccountB.CheckBalance())
	fmt.Println("Making joint account...")
	testJointAccount, _ = bank.JointAccount(&newAccountA, &newAccountB)
	fmt.Println("After joining accounts A and B:")
	fmt.Println("Account A Balance: $", newAccountA.CheckBalance())
	fmt.Println("Account B Balance: $", newAccountB.CheckBalance())
	fmt.Println("→ Joint Account Balance: $", testJointAccount.CheckBalance())
}
