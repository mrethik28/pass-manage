package input

import "fmt"

func GetMasterKey() string {
	var key string
	fmt.Println("Please enter the master key to access all saved passwords")
	fmt.Scanf("%s", &key)
	return key
}
