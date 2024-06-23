package core

import (
	"fmt"
	"github.com/voodooEntity/go-clibuddy/src/conversation"
	"github.com/voodooEntity/go-clibuddy/src/envinfo"
	"github.com/voodooEntity/go-clibuddy/src/ollamapi"
	"github.com/voodooEntity/go-clibuddy/src/promptbuilder"
	"os"
)

type Core struct {
	Api           *ollamapi.OllamApi
	Conversation  *conversation.Conversation
	Environment   *envinfo.EnvInfo
	PromptBuilder *promptbuilder.PromptBuilder
}

func New() *Core {
	ei := envinfo.New()
	c := Core{
		Api:           ollamapi.New("http://localhost:11434/api/generate", "codellama"),
		Conversation:  conversation.New(),
		Environment:   ei,
		PromptBuilder: promptbuilder.New(ei),
	}

	if "" == c.Conversation.Argv1 {
		fmt.Println("No task/question given.")
		os.Exit(0)
	}
	return &c
}

func (c *Core) Execute() {

	prompt := c.PromptBuilder.BuildPrompt(c.Conversation.Argv1)

	// Ask the question using the API
	response, err := c.Api.Ask(prompt)
	if err != nil {
		fmt.Println("Error asking question:", err)
		return
	}

	// Output the response
	fmt.Println("Response:\n", response)
}
