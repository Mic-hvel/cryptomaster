package main

import "coin.com/go/crypto/api"

func main () {
	rate, err := api.GetRate("BTC")

	print(rate, err)
}