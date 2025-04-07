package modules

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var verbose bool

// SetVerbose sets the verbose mode for the package
func SetVerbose(v bool) {
	verbose = v
}

// CheckContainerEscape checks for potential container/VM escape vectors
func CheckContainerEscape() error {
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("%s %s\n", yellow("[*]"), "Checking for container/VM escape opportunities...")

	// Check if we're in a container
	cgroupContent, err := os.ReadFile("/proc/1/cgroup")
	if err != nil {
		if verbose {
			fmt.Printf("%s Could not read /proc/1/cgroup: %v\n", red("[-]"), err)
		}
		return err
	}

	content := string(cgroupContent)
	if strings.Contains(content, "docker") {
		fmt.Printf("%s %s\n", yellow("[*]"), "Running inside Docker container")
		checkDockerEscape()
	} else if strings.Contains(content, "lxc") {
		fmt.Printf("%s %s\n", yellow("[*]"), "Running inside LXC container")
	} else {
		fmt.Printf("%s %s\n", yellow("[!]"), "No container environment detected")
	}

	// Check for common mounted docker socket
	if _, err := os.Stat("/var/run/docker.sock"); err != nil {
		if verbose {
			fmt.Printf("%s Docker socket check: %v\n", red("[-]"), err)
		}
	} else {
		fmt.Printf("%s %s\n", yellow("[*]"), "Docker socket is mounted and accessible!")
		fmt.Printf("%s %s\n", yellow("[*]"), "Potential escape vector through Docker API")
	}

	// Check for privileged mode
	capContent, err := os.ReadFile("/proc/1/status")
	if err != nil {
		if verbose {
			fmt.Printf("%s Could not read process capabilities: %v\n", red("[-]"), err)
		}
	} else {
		if strings.Contains(string(capContent), "CapEff:\t0000003fffffffff") {
			fmt.Printf("%s %s\n", yellow("[*]"), "Container is running in privileged mode!")
		} else {
			fmt.Printf("%s %s\n", yellow("[!]"), "Container is not running in privileged mode")
		}
	}

	// Check for mounted host paths
	mountContent, err := os.ReadFile("/proc/mounts")
	if err != nil {
		if verbose {
			fmt.Printf("%s Could not read mount points: %v\n", red("[-]"), err)
		}
	} else {
		mounts := string(mountContent)
		if strings.Contains(mounts, "/host/") {
			fmt.Printf("%s %s\n", yellow("[*]"), "Host filesystem appears to be mounted!")
		} else {
			fmt.Printf("%s %s\n", yellow("[!]"), "No host filesystem mounts detected")
		}
	}

	fmt.Printf("%s %s\n", green("[+]"), "Container escape check completed")
	return nil
}

func checkDockerEscape() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// Check for common docker escape vectors
	vectors := []struct {
		path string
		desc string
	}{
		{"/var/run/docker.sock", "Docker socket mounted"},
		{"/host", "Host filesystem mounted"},
		{"/proc/sys/kernel/core_pattern", "Writable core_pattern (container breakout)"},
	}

	foundVectors := false
	for _, v := range vectors {
		if _, err := os.Stat(v.path); err == nil {
			fmt.Printf("%s %s: %s\n", yellow("[*]"), v.desc, v.path)
			foundVectors = true
		} else if verbose {
			fmt.Printf("%s Checking %s: %v\n", red("[-]"), v.path, err)
		}
	}

	if !foundVectors {
		fmt.Printf("%s %s\n", yellow("[!]"), "No common Docker escape vectors found")
	}
}
