package ch7

import "errors"

type Address struct {
	address string
}

func NewAddress(address string) (*Address, error) {
	if !validateAddress(address) {
		return nil, errors.New("invalid address")
	}
	return &Address{address: address}, nil
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
