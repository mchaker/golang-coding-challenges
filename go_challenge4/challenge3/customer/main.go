package customer

// Customer represents a customer, intended for use by the bank package's
// bankAccount struct
type Customer struct {
	name string
}

// ChangeName changes the customer's name
func (bankCustomer *Customer) ChangeName(newCustomerName string) {
	bankCustomer.name = newCustomerName
}

// CheckName returns the Customer's name
func (bankCustomer *Customer) CheckName() string {
	return bankCustomer.name
}
