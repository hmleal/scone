package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/briandowns/spinner"
)

type Scoop struct {
	path    string
	buckets []string
}

func NewScoop() (Scoop, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return Scoop{}, errors.New("Can't find user home directory")
	}

	scoopPath := fmt.Sprintf("%s\\Scoop", userPath)

	if _, err := os.Stat(scoopPath); os.IsNotExist(err) {
		return Scoop{}, errors.New("Can't find scoop directory")
	}

	return Scoop{
		path: scoopPath,
		buckets: []string{
			"buckets\\main",
			"buckets\\nerd-fonts",
			"buckets\\games",
		},
	}, nil
}

func RunCmd(value string) {
	cmd := exec.Command("powershell", "git", "pull")
	cmd.Dir = value

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
}

func RunCommands(s *Scoop) {
	wg := sync.WaitGroup{}

	for _, value := range s.buckets {
		wg.Add(1)
		value := value
		go func() {
			defer wg.Done()
			RunCmd(fmt.Sprintf("%s\\%s", s.path, value))
		}()
	}

	wg.Wait()
}

func main() {
	scoop, _ := NewScoop()

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " thinking..."
	s.FinalMSG = "Scone was update successfully!"
	s.Color("blue")

	start := time.Now()
	s.Start()
	RunCommands(&scoop)
	s.Stop()
	fmt.Printf("\nTotal of seconds: %v", time.Since(start))
}
