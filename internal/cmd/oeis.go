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

type cmdN struct {
	val int
	ind int
	err error
}

func runCmdWithN(command string, n int, reschan chan cmdN) int {
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

	reschan <- cmdN{val, n, err}

	return val
}

func OeisCmdFn(cmd string, start int, end int) error {
	lis := make([]int, end)

	reschan := make(chan cmdN)

	for i := start; i <= end; i++ {
		go runCmdWithN(cmd, i, reschan)
	}

	for i := start; i <= end; i++ {
		res := <-reschan
		lis[res.ind-start] = res.val
	}

	url := "https://oeis.org/search?q="

	for _, num := range lis {
		url += strconv.Itoa(num) + "%2C"
	}

	utils.BrowserOpen(url)

	return nil
}
