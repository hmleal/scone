package scoop

import (
	"errors"
	"os"
	"path/filepath"
)

type Scoop struct {
	Path    string
	Buckets []string
}

type Command struct {
	Options   string
	Directory string
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

func (s *Scoop) UpdateBuckets() []Command {
	cmds := []Command{}

	for _, b := range s.Buckets {
		cmds = append(cmds, Command{Options: "git pull", Directory: filepath.Join(s.Path, b)})
	}

	return cmds
}
