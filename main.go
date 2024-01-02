package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hmleal/scone/scoop"
)

func RunCommand(command scoop.Command) {
	cmd := exec.Command("powershell", command.Options)
	cmd.Dir = command.Directory

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
}

func RunCommands(cmds []scoop.Command) {
	wg := sync.WaitGroup{}

	for _, cmd := range cmds {
		wg.Add(1)
		cmd := cmd
		go func() {
			defer wg.Done()
			RunCommand(cmd)
		}()
	}

	wg.Wait()
}

func main() {
	scoop, _ := scoop.NewScoop()

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Updating Scone..."
	s.FinalMSG = "Scone was update successfully!"
	s.Color("blue")

	start := time.Now()
	s.Start()
	RunCommands(scoop.UpdateBuckets())
	s.Stop()
	fmt.Printf("\n\nTotal of seconds: %v\n\n", time.Since(start))
}
