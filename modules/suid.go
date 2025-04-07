package modules 

import (
	"fmt"
	"os/exec"
)

func CheckSUID() {
	fmt.Println("[*] Searching for SUID/SGID files:")
	cmd := exec.Command("find", "/", "-perm", "6000", "-type", "f", "-exec", "ls", "-la", "{}", ";")
	out, err := cmd.CombinedOutput()
	if err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println("[-] Error:", err)
	}
}
