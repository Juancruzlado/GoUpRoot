package modules

import (
	"fmt"
	"os"
)

func CheckSensitiveFilePerms() {
	fmt.Println("[+] Checking permissions on /etc/passwd and /etc/shadow...")

	files := []string{"/etc/passwd", "/etc/shadow"}

	for _, f := range files {
		info, err := os.Stat(f)
		if err != nil {
			fmt.Printf("  [-] Cannot access %s: %v\n", f, err)
			continue
		}

		mode := info.Mode().Perm()
		fmt.Printf("  [*] %s - permissions: %#o\n", f, mode)

		if mode&0o002 != 0 {
			fmt.Printf("    [!] World-writable: %s\n", f)
		}
	}
}

