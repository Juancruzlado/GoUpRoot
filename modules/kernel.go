package modules

import (
	"fmt"
	"os/exec"
)

func CheckKernel() {
	fmt.Println("[+] Kernel Info:")
	out, _ := exec.Command("uname", "-a").Output()
	fmt.Println(string(out))
}
