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

func (pb *PromptBuilder) BuildCommandPrompt(task string) string {
	return fmt.Sprintf(`Given the environment information:
%s

Provide a CLI/bash command to accomplish the following task:
%s

Your response MUST only contain the resulting command and nothing else.
Your response MUST be valid to be executed as command directly.
It is FORBIDDEN to respond with any explaination or description.
`, pb.envInfo.GetPromptFormat(), task)
}

func (pb *PromptBuilder) BuildExplanationPrompt(command string) string {
	return fmt.Sprintf(`Given the environment information:
%s

Provide a description/explaination on what the following command would do: 
%s`, pb.envInfo.GetPromptFormat(), command)
}
