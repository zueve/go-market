package billing

import "errors"

var (
	ErrInvoiceExists   = errors.New("invoice already exists")
	ErrNotEnoughtMoney = errors.New("out of money")
)
