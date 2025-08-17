package cmdFn

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
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

type cmds struct {
	current cmdN
	reschan chan cmdN
}

func (c cmds) await() cmds {
	return cmds{current: <-c.reschan, reschan: c.reschan}
}

type model struct {
	start    int
	end      int
	lis      []int
	cmd      string
	current  int
	progress progress.Model
}

func (m model) Init() tea.Cmd {
	reschan := make(chan cmdN)

	go func() {
		for i := m.start; i <= m.end; i++ {
			go runCmdWithN(m.cmd, i, reschan)
		}
	}()

	return func() tea.Msg {
		return cmds{current: <-reschan, reschan: reschan}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	case cmds:
		m.lis[msg.current.ind-m.start] = msg.current.val
		m.current++
		progressCmd := m.progress.SetPercent((float64(m.current) / float64(m.end-m.start+1)) * 100)

		if m.current < m.end-m.start+1 {
			return m, tea.Batch(progressCmd, func() tea.Msg {
				return msg.await()
			})
		} else {
			return m, tea.Quit
		}
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	return m.progress.View()
}

func OeisCmdFn(cmd string, start int, end int) error {
	p := tea.NewProgram(model{
		start:    start,
		end:      end,
		progress: progress.New(),
		cmd:      cmd,
		lis:      make([]int, end-start+1),
	})

	m, err := p.Run()
	if err != nil {
		log.Fatal()
		return err
	}

	lis := m.(model).lis

	// reschan := make(chan cmdN)
	//
	// for i := start; i <= end; i++ {
	// 	go runCmdWithN(cmd, i, reschan)
	// }
	//
	// for i := start; i <= end; i++ {
	// 	res := <-reschan
	// 	lis[res.ind-start] = res.val
	// }

	url := "https://oeis.org/search?q="

	for _, num := range lis {
		url += strconv.Itoa(num) + "%2C"
	}

	utils.BrowserOpen(url)

	return nil
}
