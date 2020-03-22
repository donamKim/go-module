package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return "HELLO WORLD!!"
}

func Hello2() string {
	return "HELLO WORLD2"
}

func Quote() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}