package core

import (
	"fmt"
	"github.com/voodooEntity/go-clibuddy/src/conversation"
	"github.com/voodooEntity/go-clibuddy/src/envinfo"
	"github.com/voodooEntity/go-clibuddy/src/executor"
	"github.com/voodooEntity/go-clibuddy/src/ollamapi"
	"github.com/voodooEntity/go-clibuddy/src/promptbuilder"
	"os"
)

type Core struct {
	Api           *ollamapi.OllamApi
	Conversation  *conversation.Conversation
	Environment   *envinfo.EnvInfo
	PromptBuilder *promptbuilder.PromptBuilder
	Executor      *executor.Executor
}

func New() *Core {
	ei := envinfo.New()

	c := Core{
		Api:           ollamapi.New("http://localhost:11434/api/generate", "codestral"), // codellama
		Conversation:  conversation.New(),
		Environment:   ei,
		PromptBuilder: promptbuilder.New(ei),
		Executor:      executor.New(ei),
	}

	if "" == c.Conversation.Argv1 {
		fmt.Println("No task/question given.")
		fmt.Println("\n")
		os.Exit(0)
	}
	return &c
}

func (c *Core) Execute() {

	prompt := c.PromptBuilder.BuildCommandPrompt(c.Conversation.Argv1)

	// Ask the question using the API
	command, err := c.Api.Ask(prompt)
	if err != nil {
		fmt.Println("Error asking question:", err)
		return
	}
	fmt.Println("The generated command is:", command)

	options := map[string]string{"explain": "Explains the generated command", "run": "Executes the generated command", "exit": "End the clibuddy conversation"}
	for {
		action := c.Conversation.ComplexQuestion("How you want to proceed?", options)
		switch action {
		case "explain":
			explainPrompt := c.PromptBuilder.BuildExplanationPrompt(command)
			explaination, err := c.Api.Ask(explainPrompt)
			if err != nil {
				fmt.Println("Error while trying to ask the llm - exiting:", err)
				os.Exit(0)
			}
			fmt.Println(explaination)
		case "run":
			result, err := c.Executor.Do(command)
			if err != nil {
				fmt.Println("Error while executing command - exiting:", err)
				os.Exit(0)
			}
			fmt.Println("Command executed successfully:")
			fmt.Println(result)
			fmt.Println("\n")
		case "exit":
			fmt.Println("\n")
			fmt.Println("So long and thanks for all the fish!")
			os.Exit(0)
		}
	}

}
