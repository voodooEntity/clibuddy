package envinfo

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type EnvInfo struct {
	CurrentDir           string
	Files                []string
	Directories          []string
	EnvVars              map[string]string
	AvailableCommands    []string
	OS                   string
	Shell                string
	GitBranch            string
	GitBranchesAvailable []string
}

func New() *EnvInfo {
	ei := &EnvInfo{}
	ei.collectInfo()
	return ei
}

func (ei *EnvInfo) GetPromptFormat() string {
	return fmt.Sprintf(`Current directory: %s
Files in current directory: [%s]
Directories in current directory: [%s]
Environment variables: {%s}
Available commands: [%s]
OS: %s
Shell: %s
Current git branch: %s
Available git branches: [%s]`,
		ei.CurrentDir,
		strings.Join(ei.Files, ", "),
		strings.Join(ei.Directories, ", "),
		ei.formatMap(ei.EnvVars),
		strings.Join(ei.AvailableCommands, ", "),
		ei.OS,
		ei.Shell,
		ei.GitBranch,
		strings.Join(ei.GitBranchesAvailable, ", "),
	)
}

func (ei *EnvInfo) collectInfo() {
	ei.CurrentDir = ei.getCurrentDir()
	ei.Files, ei.Directories = ei.getFilesAndDirs()
	ei.EnvVars = ei.getEnvVars()
	ei.AvailableCommands = ei.getAvailableCommands()
	ei.OS = ei.getOSInfo()
	ei.Shell = ei.getShell()
	ei.GitBranch = ei.getGitBranch()
	ei.GitBranchesAvailable = ei.getGitBranches()
}

func (ei *EnvInfo) getCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentDir
}

func (ei *EnvInfo) getFilesAndDirs() ([]string, []string) {
	files := []string{}
	dirs := []string{}
	entries, err := os.ReadDir(ei.CurrentDir)
	if err != nil {
		return files, dirs
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		} else {
			files = append(files, entry.Name())
		}
	}
	return files, dirs
}

func (ei *EnvInfo) getEnvVars() map[string]string {
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}
	return envVars
}

func (ei *EnvInfo) getAvailableCommands() []string {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return []string{}
	}

	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	commandSet := make(map[string]struct{})

	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, file := range files {
			if !file.IsDir() {
				commandSet[file.Name()] = struct{}{}
			}
		}
	}

	commands := make([]string, 0, len(commandSet))
	for cmd := range commandSet {
		commands = append(commands, cmd)
	}

	return commands
}

func (ei *EnvInfo) getOSInfo() string {
	return runtime.GOOS
}

func (ei *EnvInfo) getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "unknown"
	}
	return shell
}

func (ei *EnvInfo) getGitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = ei.CurrentDir
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func (ei *EnvInfo) getGitBranches() []string {
	cmd := exec.Command("git", "branch", "--list")
	cmd.Dir = ei.CurrentDir
	output, err := cmd.Output()
	if err != nil {
		return []string{}
	}
	branches := strings.Split(string(output), "\n")
	for i, branch := range branches {
		branches[i] = strings.TrimSpace(branch)
	}
	return branches
}

func (ei *EnvInfo) formatMap(m map[string]string) string {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%s: %s, ", k, v))
	}
	s := sb.String()
	if len(s) > 2 {
		s = s[:len(s)-2] // Remove the trailing comma and space
	}
	return s
}
