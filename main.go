package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

	scoopPath := filepath.Join(userPath, "Scoop")

	if _, err := os.Stat(scoopPath); os.IsNotExist(err) {
		return Scoop{}, errors.New("Can't find scoop directory")
	}

	buckets, err := os.ReadDir(filepath.Join(scoopPath, "buckets"))
	if err != nil {
		return Scoop{}, errors.New("Can't find buckets directory")
	}

	var enabledBuckets []string
	for _, b := range buckets {
		if !b.IsDir() {
			continue
		}
		enabledBuckets = append(enabledBuckets, filepath.Join("buckets", b.Name()))
	}

	return Scoop{
		path:    scoopPath,
		buckets: enabledBuckets,
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
			RunCmd(filepath.Join(s.path, value))
		}()
	}

	wg.Wait()
}

func main() {
	scoop, _ := NewScoop()

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Updating Scone..."
	s.FinalMSG = "Scone was update successfully!"
	s.Color("blue")

	start := time.Now()
	s.Start()
	RunCommands(&scoop)
	s.Stop()
	fmt.Printf("\n\nTotal of seconds: %v\n\n", time.Since(start))
}
