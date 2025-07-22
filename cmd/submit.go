package submit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
		if !file.IsDir() {
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

	fmt.Println(file)

	os.Chdir(prevdir)
	return nil
}
