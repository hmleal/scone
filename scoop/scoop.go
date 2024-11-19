package scoop

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/animations"
)

type Scoop struct {
	Path    string
	Buckets []string
}

type Bucket struct {
	Name      string
	Source    string
	Manifests int
}

type Command struct {
	Title     string
	Options   string
	Directory string
}

func RunCommand(command Command, s *ysmrr.Spinner) {
	cmd := exec.Command("powershell", command.Options)
	cmd.Dir = command.Directory

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}

	s.Complete()
}

func RunCommands(cmds []Command) {
	// The extras bucket was removed successfully.
	// The extras bucket was added successfully.
	sm := ysmrr.NewSpinnerManager(ysmrr.WithAnimation(animations.Arc))

	sm.Start()
	defer sm.Stop()

	wg := sync.WaitGroup{}

	for _, cmd := range cmds {
		wg.Add(1)
		cmd := cmd
		go func() {
			defer wg.Done()
			RunCommand(cmd, sm.AddSpinner(cmd.Title))
		}()
	}

	wg.Wait()
}

func NewScoop() (*Scoop, error) {
	scoop := Scoop{}

	userPath, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("can't find user home directory")
	}

	scoopPath := filepath.Join(userPath, "Scoop")

	if _, err := os.Stat(scoopPath); os.IsNotExist(err) {
		return nil, errors.New("can't find scoop directory")
	}

	buckets, err := os.ReadDir(filepath.Join(scoopPath, "buckets"))
	if err != nil {
		return nil, errors.New("can't find buckets directory")
	}

	var enabledBuckets []string
	for _, b := range buckets {
		if !b.IsDir() {
			continue
		}
		enabledBuckets = append(enabledBuckets, filepath.Join("buckets", b.Name()))
	}

	scoop.Path = scoopPath
	scoop.Buckets = enabledBuckets

	return &scoop, nil
}

func (s *Scoop) UpdateBuckets() {
	cmds := []Command{}

	for _, b := range s.Buckets {
		cmds = append(cmds, Command{Title: fmt.Sprintf("Updating %s...", b), Options: "git pull", Directory: filepath.Join(s.Path, b)})
	}

	start := time.Now()
	RunCommands(cmds)
	fmt.Printf("\nTotal of seconds: %v\n\n", time.Since(start))
}

func (s *Scoop) AddBucket(name string) error {
	bucketPath := filepath.Join(s.Path, "buckets")

	cmds := []Command{}
	cmds = append(cmds, Command{
		Title:     fmt.Sprintf("Adding %s bucket", name),
		Options:   fmt.Sprintf("git clone https://github.com/ScoopInstaller/%s", name),
		Directory: bucketPath,
	})

	start := time.Now()
	RunCommands(cmds)
	fmt.Printf("\n\nTotal of seconds: %v\n\n", time.Since(start))

	return nil
}

func (s *Scoop) RemoveBucket(name string) error {
	bucketPath := filepath.Join(s.Path, "buckets", name)

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return errors.New("can't find bucket directory")
	}

	os.RemoveAll(bucketPath)

	return nil
}

// scoop.Add
// scoop.Delete
// scoop.Update
// scoop.List

// scoop.Buckets.Add
// scoop.Buckets.Remove
// scoop.Buckets.Update
// scoop.Buckets.List
