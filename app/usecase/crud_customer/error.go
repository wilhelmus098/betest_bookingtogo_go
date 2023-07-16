package crud_customer

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrCustomerNotFound = errors.New("Customer not found")
