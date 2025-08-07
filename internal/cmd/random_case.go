package cmdFn

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func caseGenerator(generateCmd string, solveCmd string, file_index int, dir string, reschan chan error) {
	// 存在するなら削除
	input_path := path.Join("./", dir, fmt.Sprintf("random_case_%d.in", file_index))
	out_path := path.Join("./", dir, fmt.Sprintf("random_case_%d.out", file_index))

	exec.Command("rm", input_path).Run()
	exec.Command("rm", out_path).Run()

	shell := os.Getenv("SHELL")

	input_generate_cmd := exec.Command(shell, "-c", generateCmd)
	out, err := input_generate_cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
		reschan <- err
		return
	}

	input := string(out)
	err = os.WriteFile(input_path, []byte(input), 0644)

	if err != nil {
		log.Fatal(err)
		reschan <- err
		return
	}

	output_solve_cmd := exec.Command(shell, "-c", solveCmd)
	output_solve_cmd.Stdin = strings.NewReader(input)

	out, err = output_solve_cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
		reschan <- err
		return
	}

	err = os.WriteFile(out_path, []byte(out), 0644)

	if err != nil {
		log.Fatal(err)
		reschan <- err
		return
	}

	reschan <- nil
}

func RandomCaseFn(generateCmd string, solveCmd string, n int, dir string) error {
	c := make(chan error, n)
	for i := 0; i < n; i++ {
		go caseGenerator(generateCmd, solveCmd, i, dir, c)
	}

	for i := 0; i < n; i++ {
		err := <-c

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
