package modules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckPasswords() {
	fmt.Println("[+] Searching for passwords in config files (/etc)...")
	filepath.Walk("/etc", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".conf") || strings.Contains(path, "config") {
			data, err := os.ReadFile(path)
			if err == nil {
				content := string(data)
				if strings.Contains(content, "password") || strings.Contains(content, "secret") {
					fmt.Printf("  [!] Possible secret in: %s\n", path)
				}
			}
		}
		return nil
	})
}

