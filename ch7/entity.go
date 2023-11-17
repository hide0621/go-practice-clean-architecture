package ch7

import "errors"

type Address struct {
	address string
}

func NewAddress(address string) *Address {
	if !validateAddress(address) {
		panic(errors.New("invalid address"))
	}
	return &Address{address: address}
}

func validateAddress(address string) bool {
	// Implement address validation logic
	return true
}

type Mail struct {
	id          int
	ownAddress  *Address
	toAddresses *[]Address
	mainText    string
}
