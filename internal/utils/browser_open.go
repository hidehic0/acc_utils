package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func getOpenCmd() string {
	switch runtime.GOOS {
	case "darwin":
		return "open"
	case "linux":
		return "xdg-open"
	case "freebsd":
		return "xdg-open"
	default:
		return ""
	}

}

func BrowserOpen(url string) error {
	err := exec.Command(getOpenCmd(), url).Run()

	if err != nil {
		log.Fatal(err)
		os.Exit(256)
	}

	return nil
}
