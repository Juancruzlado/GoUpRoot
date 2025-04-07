package modules

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func CheckProcesses() {
	fmt.Println("[+] Listing processes running as root:")

	procFiles, _ := ioutil.ReadDir("/proc")
	for _, f := range procFiles {
		if f.IsDir() && isNumeric(f.Name()) {
			statusPath := "/proc/" + f.Name() + "/status"
			data, err := ioutil.ReadFile(statusPath)
			if err == nil && strings.Contains(string(data), "Uid:\t0") {
				cmdline, _ := ioutil.ReadFile("/proc/" + f.Name() + "/cmdline")
				fmt.Printf("  [#] PID %s is root process: %s\n", f.Name(), strings.ReplaceAll(string(cmdline), "\x00", " "))
			}
		}
	}
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

