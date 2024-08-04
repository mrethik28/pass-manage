package main

import (
	"fmt"
	"main/input"
	"main/validate"
)

func main() {
	fmt.Println("Welcome to your local password manager 🙏")
	masterkey := input.GetMasterKey()
	fmt.Println(validate.ValidateMasterKey(masterkey))
}
