package hello

import "rsc.io/quote/v3"

func Hello() string {
	// return quote.HelloV3()
	return Run()
}

func Proverb() string {
	return quote.Concurrency()
}
