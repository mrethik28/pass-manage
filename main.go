package main

import (
	"fmt"
	interaction "main/io"
	"main/validate"
)

func main() {
	fmt.Println("Welcome to your local password manager 🙏")
	masterkey := interaction.GetMasterKey()
	fmt.Println(validate.ValidateMasterKey(masterkey))
	interaction.SelectPasword()

}
