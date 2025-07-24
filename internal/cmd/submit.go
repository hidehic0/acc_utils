package submit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"hidehic0/acc_utils/internal/utils"

	"github.com/koki-develop/go-fzf"
)

func SubmitFn(dir string) error {
	files, err := os.ReadDir("./" + dir)

	if err != nil {
		log.Fatal(err)
		return err
	}

	items := []string{}

	for _, file := range files {
		_, ok := utils.GetFileConfig()[file.Name()]
		if !file.IsDir() && ok {
			items = append(items, file.Name())
		}
	}

	f, err := fzf.New()

	if err != nil {
		log.Fatal(err)
		return err
	}

	inds, err := f.Find(items, func(i int) string { return items[i] })

	if err != nil {
		log.Fatal(err)
		return err
	}

	if len(inds) == 0 {
		fmt.Println("Item not selected")
		return nil
	}
	file := items[inds[0]]
	prevdir, _ := filepath.Abs(".")
	os.Chdir("./" + dir)

	var cmd string

	if val, ok := utils.GetFileConfig()[file]; ok {
		cmd = val.Cmd
	} else {
		fmt.Printf("%s config not found\n", file)
		os.Exit(256)
	}

	shell := os.Getenv("SHELL")
	shell_cmd := exec.Command(shell, "-ic", cmd)
	err = shell_cmd.Start()

	if err != nil {
		log.Fatal(err)
		return err
	}

	os.Chdir(prevdir)

	// open browser
	utils.BrowserOpen(utils.GetTaskInfomation()[dir].Url + "#editor")

	return nil
}
