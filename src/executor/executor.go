package executor

import (
	"bytes"
	"fmt"
	"github.com/voodooEntity/go-clibuddy/src/envinfo"
	"os/exec"
)

// Executor struct holds the environment information
type Executor struct {
	Env *envinfo.EnvInfo
}

// New creates a new instance of Executor with provided EnvInfo
func New(env *envinfo.EnvInfo) *Executor {
	return &Executor{
		Env: env,
	}
}

// Do executes the given command string in the current directory
func (e *Executor) Do(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = e.Env.CurrentDir

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v\n%s", err, stderr.String())
	}

	return stdout.String(), nil
}
