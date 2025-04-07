package modules

import (
	"fmt"
	"os"
)

func CheckCron() {
	fmt.Println("[*] Checking for cronjobs:")
	files := []string{"/etc/crontab", "/etc/cron.d", "/var/spool/cron", "/etc/cron.hourly"}
	for _, path := range files {
			fmt.Println(" -", path)
			info, err := os.ReadFile(path)
			if err == nil {
					fmt.Println(string(info))
			}
	}
}

