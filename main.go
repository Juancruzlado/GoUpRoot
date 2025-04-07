package main

import (
	"fmt"

	"github.com/juan/localenum/modules"
)

func printBanner() {
	banner := `
   ______      __  __      ____              __ 
  / ____/___  / / / /___  / __ \____  ____  / /_
 / / __/ __ \/ / / / __ \/ /_/ / __ \/ __ \/ __/
/ /_/ / /_/ / /_/ / /_/ / _, _/ /_/ / /_/ / /_  
\____/\____/\____/ .___/_/ |_|\____/\____/\__/  
                /_/                             

        GoUpRoot - Local PrivEsc Scanner
`
	fmt.Println(banner)
}

func main() {
	printBanner()

	fmt.Println("[*] Starting GoUpRoot local privilege enum...\n")

	fmt.Println("[+] Basic System Info")
	modules.CheckKernel()
	modules.CheckUsers()
	modules.CheckEnv()

	fmt.Println("[+] File & Permission Checks")
	modules.CheckSUID()
	modules.CheckCron()
	modules.CheckSensitiveFilePerms()

	fmt.Println("[+] Secrets & Credentials")
	modules.CheckPasswords()
	modules.CheckCreds()

	fmt.Println("[+] Process Discovery")
	modules.CheckProcesses()

	fmt.Println("\n[*] GoUpRoot finished. Review your findings.")
}
