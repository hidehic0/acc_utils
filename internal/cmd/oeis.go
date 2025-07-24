package cmdFn

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"hidehic0/acc_utils/internal/utils"
)

func runCmdWithN(command string, n int) int {
	shell := os.Getenv("SHELL")
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%d\n", n))

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
		os.Exit(256)
	}

	val, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {

		log.Fatal(err)
		os.Exit(256)
	}

	return val
}

func OeisCmdFn(cmd string, start int, end int) error {
	var lis []int

	for i := start; i <= end; i++ {
		lis = append(lis, runCmdWithN(cmd, i))
	}

	url := "https://oeis.org/search?q="

	for _, num := range lis {
		url += strconv.Itoa(num) + "%2C"
	}

	utils.BrowserOpen(url)

	return nil
}
