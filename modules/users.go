package modules 

import (
	"fmt"
	"os/user"
)

func CheckUsers() {
	fmt.Println("[+] Current user info:")
	u, _ := user.Current()
	fmt.Printf(" - Username: %s\n - Uid: %s\n - HomeDir: %s\n", u.Username, u.Uid, u.HomeDir)
}
