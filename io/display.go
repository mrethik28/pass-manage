package interaction

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
)

var PasswordStore = map[string]string{
	"Facebook":  "facebook123",
	"Twitter":   "twitter456",
	"Instagram": "insta789",
	"LinkedIn":  "linkedin321",
	"Snapchat":  "snapchat654",
	"TikTok":    "tiktok987",
	"Pinterest": "pinterest111",
	"Reddit":    "reddit222",
}

func SelectPasword() {
	fmt.Println("Welcome to your local password manager 🙏")
	sites := make([]string, 0, len(PasswordStore))
	for site := range PasswordStore {
		sites = append(sites, site)
	}
	prompt := promptui.Select{
		Label: "Select a site to copy password",
		Items: sites,
		Size:  8,
	}
	_, selectedSite, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	password := PasswordStore[selectedSite]

	err = clipboard.WriteAll(password)
	if err != nil {
		fmt.Printf("Failed to copy password to clipboard: %v\n", err)
		return
	}
	fmt.Printf("Password for %s copied to clipboard: %s\n", selectedSite, password)
}
