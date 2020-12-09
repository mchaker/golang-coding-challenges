package bank

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type balance struct {
	dollars int64
	cents   int64
}

// BankAccount represents a bank account
//   name: name of the account owner
//   balance: the account's balance
type bankAccount struct {
	name    string
	balance balance
}

// roundInt is based on https://web.archive.org/web/20201027140219/https://www.cockroachlabs.com/blog/rounding-implementations-in-go/
func roundInt(original int64, sigFigs int64) int64 {
	if sigFigs >= int64(len(strconv.Itoa(int(original)))) {
		return original
	}

	if math.IsNaN(float64(original)) {
		return original
	}

	if original == 0 {
		return original
	}

	roundFn := math.Ceil

	if math.Signbit(float64(original)) {
		roundFn = math.Floor
	}

	// all this int64 and int back and forth typecastying is crazy... there
	// has to be a better way
	digitToRound, err := strconv.ParseInt(fmt.Sprint(original)[sigFigs:sigFigs+1], 10, 8)

	if err != nil {
		return original
	}

	numZeroes := int64(len(fmt.Sprint(original))) - sigFigs
	tens := int64(math.Pow10(int(numZeroes)))

	rounded := float64(original) / float64(tens)

	if digitToRound >= 5 {
		rounded = roundFn(float64(original) / float64(tens))
	}

	truncated := int64(rounded)

	return truncated * tens
}

func parseDollars(amount string) (balance, error) {
	parsedAmount := balance{}

	// nothing to parse, return 0 value
	if amount == "" {
		return balance{0, 0}, nil
	}

	// Separate the dollars and cents values on the decimal point
	splitAmount := strings.Split(amount, ".")

	// This can probably be improved. There has to be a better way to signal
	// that the input was invalid than returning an empty balance
	if len(splitAmount) > 2 {
		return balance{0, 0}, fmt.Errorf("currency parsing failed: more than one decimal point")
	}

	// handle cents
	// if there is a value after the decimal point
	if len(splitAmount) > 1 {
		centsString := splitAmount[1]

		// if parsing the cents value succeeds
		if centsTest, err := strconv.ParseInt(centsString, 10, 64); err == nil {
			switch len(centsString) {
			case 1:
				parsedAmount.cents = centsTest * 10

			case 2:
				parsedAmount.cents = centsTest

			default:
				roundingDigit, _ := strconv.ParseInt(string(centsString[2]), 10, 64)
				roundingAmount := int64(0)

				if roundingDigit >= 5 {
					roundingAmount = int64(1)
				}

				tempCents, _ := strconv.ParseInt(centsString[:2], 10, 64)
				parsedAmount.cents = tempCents + roundingAmount
			}
		} else {
			if len(centsString) == 0 {
				parsedAmount.cents = 0
			} else {
				return balance{0, 0}, fmt.Errorf("currency parsing failed: parsing cents value failed")
			}
		}
	}

	// handle dollars
	dollars := int64(0)

	// if there is a dollar value to parse, parse it
	if len(splitAmount) > 0 {
		if dollarsTest, err := strconv.ParseInt(splitAmount[0], 10, 64); err == nil {
			dollars = dollarsTest
		}

		// if we have 100 cents, that is a dollar.
		// reset cents to 0 and increment the dollar amount
		if parsedAmount.cents == 100 {
			parsedAmount.dollars = dollars + 1
			parsedAmount.cents = 0
		} else {
			parsedAmount.dollars = dollars
		}
	}

	// if the input value is negative, apply the sign correctly
	if (amount[0] == '-') && (parsedAmount.cents > 0) {
		parsedAmount.cents *= -1
	}

	return parsedAmount, nil
}

func (bal balance) String() string {
	var currentBalance string

	if bal.cents < 0 {
		if bal.dollars == 0 {
			currentBalance = fmt.Sprintf("-%d.%d", bal.dollars, (bal.cents * -1))
		} else {
			currentBalance = fmt.Sprintf("%d.%d", bal.dollars, (bal.cents * -1))
		}
	} else {
		currentBalance = fmt.Sprintf("%d.%d", bal.dollars, bal.cents)
	}

	return currentBalance
}

// ===== Challenge 2 functions =====

// MakeAccount makes and returns a bankAccount
func MakeAccount(name string, initialBalance string) (bankAccount, error) {
	account := bankAccount{}

	account.name = name

	// easy input: replace blank input with zero
	if initialBalance == "" {
		initialBalance = "0"
	}

	if _, err := parseDollars(initialBalance); err != nil {
		// parsing was unsuccessful, don't increment the account balance and just return
		fmt.Println("[ERROR] Could not set initial account balance to ", initialBalance)
		fmt.Println("[ERROR] Setting balance to 0.00")
		fmt.Println("[ERROR] Error details: ", err)
	} else {
		// if parsing was successful, set the initial balance
		if initialBalance[0] == '-' {
			fmt.Println("[WARNING] Attempting to open account for ", account.name, " with a negative initial balance of ", initialBalance)
			fmt.Println("[WARNING] Since it is mean to start customers off in debt, the bank will start the account with $0.00")
		}

		account.Deposit(initialBalance)
	}

	return account, nil
}

// JointAccount makes and returns a bank account composed of the sum of
// two bankAccounts' balances and the concatenation of the names on the accounts,
// separated by an ampersand symbol.
func JointAccount(account1 *bankAccount, account2 *bankAccount) (bankAccount, error) {
	jointAccount := bankAccount{}

	jointAccount.name = fmt.Sprint("JOINT: ", account1.name, " & ", account2.name)

	// merge the first account into the joint account
	jointAccount.Deposit(account1.CheckBalance())
	account1.Withdraw(account1.CheckBalance())

	// merge the second account into the joint account
	jointAccount.Deposit(account2.CheckBalance())
	account2.Withdraw(account2.CheckBalance())

	return jointAccount, nil
}

// Deposit increases the balance in the bankAccount by the given amount
func (account *bankAccount) Deposit(amount string) {
	// if parsing was successful, increment the account balance
	if parsedAmount, err := parseDollars(amount); err != nil {
		// parsing was unsuccessful, don't increment the account balance and just return
		fmt.Println("[ERROR] Deposit unsuccessful.")
		fmt.Println("[ERROR] Error details: ", err)
	} else {
		if (parsedAmount.dollars < 0) || (parsedAmount.cents < 0) {
			err := fmt.Errorf("invalid amount (%s) and unhappy bank teller: amount to deposit must be a positive number", amount)
			fmt.Println("[ERROR] Deposit unsuccessful.")
			fmt.Println("[ERROR] Error details: ", err)
		} else {
			newCents := int64(account.balance.cents + parsedAmount.cents)
			carryDollar := int64(0)

			if newCents > 100 {
				carryDollar = 1
				newCents -= 100
			}

			account.balance.dollars += (parsedAmount.dollars + carryDollar)
			account.balance.cents = newCents
		}
	}

	return
}

// Withdraw reduces the balance in the bankAccount by the given amount
// QUESTION: Why does *bankAccount work, but bankAccount not work?
// What I mean is: the balance does not update in the object when using
// the "bankAccount" type instead of "*bankAccount". :(
func (account *bankAccount) Withdraw(amount string) {
	parsedAmount, err := parseDollars(amount)

	if err != nil {
		// parsing was unsuccessful, don't decrement the account balance and just return
		fmt.Println("[ERROR] Withdrawal unsuccessful.")
		fmt.Println("[ERROR] Error details: ", err)

		return
	}

	// if the withdraw request is for a negative amount, print an error
	if (parsedAmount.dollars < 0) || (parsedAmount.cents < 0) {
		err := fmt.Errorf("invalid amount (%s) and unhappy bank teller: amount to withdraw must be a positive number", amount)
		fmt.Println("[ERROR] Withdrawal unsuccessful.")
		fmt.Println("[ERROR] Error details: ", err)

		return
	}

	// if parsing was successful, decrement the account balance
	totalInAccount := (account.balance.dollars * 100) + account.balance.cents
	totalToWithdraw := (parsedAmount.dollars * 100) + parsedAmount.cents

	if totalToWithdraw > totalInAccount {
		err := fmt.Errorf("amount to withdraw (%s) exceeds available funds in account: %s", amount, account.name)
		fmt.Println("[ERROR] Withdrawal unsuccessful.")
		fmt.Println("[ERROR] Error details: ", err)
	} else {
		newCents := int64(account.balance.cents - parsedAmount.cents)
		carryDollar := int64(0)

		if newCents < 0 {
			carryDollar = 1
			newCents += 100
		}

		account.balance.dollars -= (parsedAmount.dollars + carryDollar)
		account.balance.cents = newCents
	}

	return
}

// CheckBalance returns the current balance in the bankAccount
func (account bankAccount) CheckBalance() string {

	return fmt.Sprint(account.balance)
}
