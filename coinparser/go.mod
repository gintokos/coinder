module github.com/gintokos/coinder/coinparser

go 1.23.2

replace github.com/gintokos/coinder/backend => ../backend

require (
	github.com/gintokos/coinder/backend v0.0.0-00010101000000-000000000000
	github.com/shopspring/decimal v1.4.0
)
