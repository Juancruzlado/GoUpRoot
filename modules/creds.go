package modules

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckCreds() {
	fmt.Println("[+] Looking for SSH keys and AWS credentials...")

	homeDirs := []string{"/root", "/home"}
	for _, base := range homeDirs {
		filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}

			// SSH keys
			if filepath.Base(path) == "id_rsa" || filepath.Base(path) == "id_ed25519" {
				fmt.Printf("  [!] Found private SSH key: %s\n", path)
			}

			// AWS creds
			if filepath.Base(path) == "credentials" && filepath.Dir(path) == base+"/.aws" {
				fmt.Printf("  [!] Found AWS credentials: %s\n", path)
			}

			return nil
		})
	}
}

