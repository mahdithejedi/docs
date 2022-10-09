package main

import (
	"SeprateChain/helpers"
)

func main() {
	hash := helpers.HashConfig{
		HashSize: 16,
	}
	hash.Add(helpers.StrHash{Data: "Ali"})
	hash.Add(helpers.StrHash{Data: "Mahdi"})
	hash.Add(helpers.StrHash{Data: "Mahdi"})
	hash.Add(helpers.StrHash{Data: "Mahdi"})
	hash.Add(helpers.StrHash{Data: "Mahdi"})
	hash.Add(helpers.StrHash{Data: "Mahdi"})
	hash.View()

}
