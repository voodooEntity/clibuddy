package promptbuilder

import (
	"fmt"
	"github.com/voodooEntity/go-clibuddy/src/envinfo"
)

type PromptBuilder struct {
	envInfo *envinfo.EnvInfo
}

// New creates a new instance of PromptBuilder and initializes it with an EnvInfo instance
func New(envInfo *envinfo.EnvInfo) *PromptBuilder {
	return &PromptBuilder{
		envInfo: envInfo,
	}
}

// BuildPrompt constructs a prompt for the LLM to generate a CLI/bash command for a given task
func (pb *PromptBuilder) BuildPrompt2(task string) string {
	envInfoFormat := pb.envInfo.GetPromptFormat()
	return fmt.Sprintf(`Given the following environment information:
%s

Write a CLI/bash command that solves the following task:
%s

The command must be printed cleanly with no additional description, quotes, or extra characters. It should be ready to execute directly in the given environment.`, envInfoFormat, task)
}

func (pb *PromptBuilder) BuildPrompt(task string) string {
	return fmt.Sprintf(`Given the environment information:
%s

Provide a CLI/bash command to accomplish the following task:
%s

Your response MUST only contain the resulting command and nothing else.
Your response MUST be valid to be executed as command directly.
It is FORBIDDEN to respond with any explaination or description.
`, pb.envInfo.GetPromptFormat(), task)
}
