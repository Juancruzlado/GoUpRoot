package modules

import (
	"fmt"
	"os"
)

func CheckEnv() {
	fmt.Println("[+] Environment variables:")
	for _, e := range os.Environ() {
			fmt.Println(" -", e)
	}
}
