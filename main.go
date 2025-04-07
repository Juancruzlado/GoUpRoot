package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/juan/localenum/modules"
)

// Results stores the enumeration results for JSON output
type Results struct {
	SystemInfo struct {
		Kernel     string   `json:"kernel"`
		Users      []string `json:"users"`
		EnvVars    []string `json:"env_vars"`
	} `json:"system_info"`
	Vulnerabilities struct {
		SUID        []string `json:"suid_files"`
		Cron        []string `json:"cron_jobs"`
		Permissions []string `json:"dangerous_permissions"`
		Secrets     []string `json:"exposed_secrets"`
		Container   []string `json:"container_vectors"`
	} `json:"vulnerabilities"`
	Errors []string `json:"errors,omitempty"`
}

var (
	verbose    bool
	jsonOutput bool
	results    Results
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output (shorthand)")
	flag.BoolVar(&jsonOutput, "json", false, "Output results in JSON format")
	flag.Parse()

	// Set verbose mode in modules package
	modules.SetVerbose(verbose)
}

func printBanner() {
	cyan := color.New(color.FgCyan).SprintFunc()
	banner := cyan(`
   ______      __  __      ____              __ 
  / ____/___  / / / /___  / __ \____  ____  / /_
 / / __/ __ \/ / / / __ \/ /_/ / __ \/ __ \/ __/
/ /_/ / /_/ / /_/ / /_/ / _, _/ /_/ / /_/ / /_  
\____/\____/\____/ .___/_/ |_|\____/\____/\__/  
                /_/                             

        GoUpRoot - Local PrivEsc Scanner ~ by xjxcxn3l
`)
	fmt.Println(banner)
}

func main() {
	printBanner()

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	
	if !jsonOutput {
		fmt.Printf("\n%s\n", yellow("[*] Starting GoUpRoot local privilege enum..."))
	}

	// Basic System Info
	if !jsonOutput {
		fmt.Printf("\n%s\n", green("[+] Basic System Info"))
	}
	
	modules.CheckKernel()
	modules.CheckUsers()
	modules.CheckEnv()

	// File & Permission Checks
	if !jsonOutput {
		fmt.Printf("\n%s\n", green("[+] File & Permission Checks"))
	}
	modules.CheckSUID()
	modules.CheckCron()
	modules.CheckSensitiveFilePerms()

	// Container/VM Escape Checks
	if !jsonOutput {
		fmt.Printf("\n%s\n", green("[+] Container/VM Escape Checks"))
	}
	if err := modules.CheckContainerEscape(); err != nil && verbose {
		fmt.Printf("%s Error checking container escape vectors: %v\n", red("[-]"), err)
	}

	// Secrets & Credentials
	if !jsonOutput {
		fmt.Printf("\n%s\n", green("[+] Secrets & Credentials"))
	}
	modules.CheckPasswords()
	modules.CheckCreds()

	// Process Discovery
	if !jsonOutput {
		fmt.Printf("\n%s\n", green("[+] Process Discovery"))
	}
	modules.CheckProcesses()

	if !jsonOutput {
		fmt.Printf("\n%s\n", yellow("[*] GoUpRoot finished. Review your findings."))
	} else {
		// Convert results to JSON and output
		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating JSON output: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	}
}
