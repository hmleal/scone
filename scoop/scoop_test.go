package scoop

import (
	"testing"
)

func TestUpdateBuckets(t *testing.T) {
	scoop := Scoop{Path: "C:\\scoop", Buckets: []string{"buckets\\main"}}
	cmds := scoop.UpdateBuckets()

	if len(cmds) != 1 {
		t.Errorf("Error")
	}

	cmd := cmds[0]

	if cmd.Options != "git pull" {
		t.Errorf("found: %s expected: %s", cmd.Options, "git pull")
	}

	if cmd.Directory != "C:\\scoop\\buckets\\main" {
		t.Errorf("found: %s expected: %s", cmd.Directory, "C:\\scoop\\buckets\\main")
	}
}
